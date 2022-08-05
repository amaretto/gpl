package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Point  int
	Expect []int
}{
	{[]int{0, 1, 2}, 1, []int{1, 2, 0}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}},
}

func TestRotate(t *testing.T) {
	for _, testCase := range testCases {
		result := rotate(testCase.Input, testCase.Point)
		if !reflect.DeepEqual(result, testCase.Expect) {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
