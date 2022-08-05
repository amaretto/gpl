package main

import "testing"

var testCases = []struct {
	Input1 string
	Input2 string
	Output bool
}{
	{"東京特許許可局許可局長", "局長許可許可局特許東京", true},
	{"part", "trap", true},
	{"tall", "tell", false},
}

func TestCommaWithBB(t *testing.T) {
	for _, testCase := range testCases {
		result := anagram(testCase.Input1, testCase.Input2)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
