package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

var testCases = []struct {
	input   string
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}{
	{"", map[rune]int{}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}, 0},
	// 1byte
	{"hoge", map[rune]int{'h': 1, 'o': 1, 'g': 1, 'e': 1}, [utf8.UTFMax + 1]int{0, 4, 0, 0, 0}, 0},
	// 2byte
	{"¢", map[rune]int{'¢': 1}, [utf8.UTFMax + 1]int{0, 0, 1, 0, 0}, 0},
	// 3byte
	{"ほげ", map[rune]int{'ほ': 1, 'げ': 1}, [utf8.UTFMax + 1]int{0, 0, 0, 2, 0}, 0},
	// 4byte
	{"𐂂𐂃𐂄", map[rune]int{'𐂂': 1, '𐂃': 1, '𐂄': 1}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 3}, 0},
	// invalid
	{"\x80", map[rune]int{}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}, 1},
}

func TestCharCount(t *testing.T) {
	for _, testCase := range testCases {
		r := bufio.NewReader(strings.NewReader(testCase.input))
		counts, utflen, invalid := charcount(*r)

		if !reflect.DeepEqual(counts, testCase.counts) || !reflect.DeepEqual(utflen, testCase.utflen) || invalid != testCase.invalid {
			t.Errorf("invalid result. testCase:%v, actual counts:%v,utflen:%v,invalid:%d", testCase, counts, utflen, invalid)
		}
	}
}
