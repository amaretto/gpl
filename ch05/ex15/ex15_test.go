package main

import (
	"testing"
)

var maxTestCases = []struct {
	Input   []int
	Output  int
	isError bool
}{
	{[]int{}, -9223372036854775808, true},
	{[]int{1}, 1, false},
	{[]int{1, 2, 3, 4}, 4, false},
	{[]int{1, 2, -3, -4}, 2, false},
	{[]int{-1, -2, -3, -4}, -1, false},
}

var minTestCases = []struct {
	Input   []int
	Output  int
	isError bool
}{
	{[]int{}, 9223372036854775807, true},
	{[]int{1}, 1, false},
	{[]int{1, 2, 3, 4}, 1, false},
	{[]int{-1, -2, 3, 4}, -2, false},
	{[]int{-1, -2, -3, -4}, -4, false},
}

func TestMax(t *testing.T) {
	for _, testCase := range maxTestCases {
		result := max(testCase.Input...)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestMaxNeedArgs(t *testing.T) {
	for _, testCase := range maxTestCases {
		result, err := maxNeedArgs(testCase.Input...)
		if testCase.isError && err == nil {
			t.Errorf("need to report error. testCaes:%v, actual:%v", testCase, result)
		}
		if !testCase.isError && err != nil {
			t.Errorf("no need to report error. testCaes:%v, actual:%v", testCase, result)
		}
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestMin(t *testing.T) {
	for _, testCase := range minTestCases {
		result := min(testCase.Input...)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestMinNeedArgs(t *testing.T) {
	for _, testCase := range minTestCases {
		result, err := minNeedArgs(testCase.Input...)
		if testCase.isError && err == nil {
			t.Errorf("need to report error. testCaes:%v, actual:%v", testCase, result)
		}
		if !testCase.isError && err != nil {
			t.Errorf("no need to report error. testCaes:%v, actual:%v", testCase, result)
		}
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
