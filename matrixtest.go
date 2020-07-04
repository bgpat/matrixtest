package matrixtest

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type kv struct {
	k string
	v interface{}
}

// Run tests with the matrixed testcases.
func Run(t *testing.T, matrix map[string]interface{}, f func(testcase interface{}) func(t *testing.T)) {
	t.Helper()

	sorted := make([]kv, 0, len(matrix))
	for k, v := range matrix {
		sorted = append(sorted, kv{k: k, v: v})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].k < sorted[j].k })

	if len(sorted) == 0 {
		return
	}

	bt, _ := typeValue(sorted[0].v)
	base := reflect.New(bt).Interface()
	t.Run("default", f(base))
	run(t, "", base, sorted, f)
}

func run(t *testing.T, prefix string, base interface{}, matrix []kv, f func(testcase interface{}) func(t *testing.T)) {
	if len(matrix) == 0 {
		return
	}

	name := matrix[0].k
	fullname := name
	if prefix != "" {
		fullname = prefix + "/" + name
	}

	testcase, err := merge(base, matrix[0].v)
	if err != nil {
		t.Fatalf("failed to merge %q: %v", name, err)
	}

	t.Run(fullname, f(testcase))
	run(t, fullname, testcase, matrix[1:], f)
	run(t, prefix, base, matrix[1:], f)
}

func merge(base, patch interface{}) (interface{}, error) {
	if patch == nil {
		return nil, errors.New("patch cannot be nil")
	}
	if base == nil {
		return patch, nil
	}

	bt, bv := typeValue(base)
	pt, pv := typeValue(patch)

	if !pt.AssignableTo(bt) {
		return nil, fmt.Errorf("%q is not assignable to %q", pt, bt)
	}

	switch bt.Kind() {
	case reflect.Struct:
		dst := reflect.New(bt).Elem()
		setFields(dst, bv, bt)
		setFields(dst, pv, pt)
		return dst.Interface(), nil
	case reflect.Map:
		dst := reflect.MakeMap(bt)
		copyMap(dst, bv)
		copyMap(dst, pv)
		return dst.Interface(), nil
	}
	return nil, fmt.Errorf("%q is not supported", bt.Kind())
}

func typeValue(i interface{}) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	return t, v
}

func setFields(dst, src reflect.Value, t reflect.Type) {
	fields := t.NumField()
	for i := 0; i < fields; i++ {
		f := t.Field(i)
		dv := dst.FieldByName(f.Name)
		if !dv.CanSet() {
			continue
		}
		sv := src.FieldByName(f.Name)
		if !sv.IsZero() {
			dv.Set(sv)
		}
	}
}

func copyMap(dst, src reflect.Value) {
	iter := src.MapRange()
	for iter.Next() {
		dst.SetMapIndex(iter.Key(), iter.Value())
	}
}
