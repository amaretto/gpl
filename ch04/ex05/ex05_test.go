package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Expect []int
}{
	{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}},
	{[]int{0, 0, 1, 2, 3}, []int{0, 1, 2, 3}},
	{[]int{0, 1, 2, 3, 3}, []int{0, 1, 2, 3}},
	{[]int{0, 0, 1, 2, 3, 3}, []int{0, 1, 2, 3}},
}

func TestRotate(t *testing.T) {
	for _, testCase := range testCases {
		result := delDup(testCase.Input)
		if !reflect.DeepEqual(result, testCase.Expect) {
			t.Errorf("invalid result. testCaes:%v", testCase)
		}
	}
}
