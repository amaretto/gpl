package main

import "testing"

var testCases = []struct {
	input  string
	expect bool
}{
	{"", true},
	{"a", true},
	{"aba", true},
	{"abc", false},
	{"aaaaaa", true},
	{"abccba", true},
	{"abcdef", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, testCase := range testCases {
		result := isPalindrome(byRune(testCase.input))
		if result != testCase.expect {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
