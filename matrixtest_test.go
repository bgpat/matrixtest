package matrixtest_test

import (
	"testing"

	"github.com/bgpat/matrixtest"
)

func TestRun(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		type testcase struct {
			Bool    bool
			Int     int
			String  string
			Pointer *struct{}
			Slice   []byte
		}
		testcases := map[string]interface{}{
			"bool":    testcase{Bool: true},
			"int":     testcase{Int: 1},
			"string":  testcase{String: "test"},
			"pointer": testcase{Pointer: &struct{}{}},
			"slice":   testcase{Slice: []byte("test")},
		}
		matrixtest.Run(t, testcases, func(testcase interface{}) func(t *testing.T) {
			return func(t *testing.T) {
				t.Log(testcase)
			}
		})
	})

	t.Run("map", func(t *testing.T) {
		matrixtest.Run(t, map[string]interface{}{
			"a": map[string]string{"a": "a"},
			"b": map[string]string{"b": "b"},
			"c": map[string]string{"c": "c"},
		}, func(testcase interface{}) func(t *testing.T) {
			return func(t *testing.T) {
				t.Log(testcase)
			}
		})
	})
}
