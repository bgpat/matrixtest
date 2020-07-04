/*
Package matrixtest provides tools for testing with matrixed testcases.

Installation:
	go get -u github.com/bgpat/matrixtest

Usage:
	func Test(t *testing.T) {
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
	}
*/
package matrixtest
