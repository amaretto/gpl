package main

import "testing"

var testCases = []struct {
	Input  string
	Output string
}{
	{"1000", "1,000"},
	{"1000000000", "1,000,000,000"},
	{"123456789", "123,456,789"},
	{"", ""},
}

func TestCommaWithBB(t *testing.T) {
	for _, testCase := range testCases {
		result := commaWithBB(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestComma(t *testing.T) {
	for _, testCase := range testCases {
		result := comma(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
