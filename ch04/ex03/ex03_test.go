package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [Length]int
	Expect [Length]int
}{
	{[Length]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, [Length]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	{[Length]int{8, 1, 3, 3, 2, 5, 4, 2, 11, 30}, [Length]int{30, 11, 2, 4, 5, 2, 3, 3, 1, 8}},
}

func TestReverse(t *testing.T) {
	for _, testCase := range testCases {
		reverse(&(testCase.Input))
		if !reflect.DeepEqual(testCase.Input, testCase.Expect) {
			t.Errorf("invalid result. testCaes:%v", testCase)
		}
	}
}
