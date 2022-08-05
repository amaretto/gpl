package main

import "testing"

var testCases = []struct {
	Input     []string
	Separator string
	Output    string
}{
	{[]string{}, "", ""},
	{[]string{"hoge"}, ",", "hoge"},
	{[]string{"hoge", "fuga"}, ",", "hoge,fuga"},
	{[]string{"hoge", "fuga", "foo"}, ",", "hoge,fuga,foo"},
	{[]string{"hoge", "fuga", "foo"}, ":", "hoge:fuga:foo"},
}

func TestJoin(t *testing.T) {
	for _, testCase := range testCases {
		result := join(testCase.Separator, testCase.Input...)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
