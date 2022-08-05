package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{"", ""},
	{"$fuga", "FUGA"},
	{"hoge", "hoge"},
	{"$hoge $fuga", "HOGE FUGA"},
	{"hoge fuga", "hoge fuga"},
	{"hoge $fuga foo $bar", "hoge FUGA foo BAR"},
}

func TestExpand(t *testing.T) {
	for _, testCase := range testCases {
		result := expand(testCase.Input, func(subs string) string {
			return strings.ToUpper(subs)
		})
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
