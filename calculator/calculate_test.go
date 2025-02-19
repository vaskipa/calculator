package calculator

import (
	"reflect"
	"testing"
)

func TestPolish(t *testing.T) {
	for _, tc := range []struct {
		expression string
		expected   []NodeData
	}{
		{"2", []NodeData{{false, 0, 2}}},
		{"2", []NodeData{{false, 0, 2}}},
		{"2 2 +", []NodeData{{false, 0, 2}, {false, 0, 2}, {true, '+', 0}}},
	} {
		t.Run(tc.expression, func(t *testing.T) {
			res, err := ToPolishNotation(tc.expression)

			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatal(res, tc.expected)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	for _, tc := range []struct {
		expression string
		expected   float64
	}{
		{"2", 2.0},
	} {
		t.Run(tc.expression, func(t *testing.T) {
			res, err := Calc(tc.expression)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatal(res, tc.expected)
			}
		})
	}
}
