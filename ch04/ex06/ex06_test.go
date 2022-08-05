package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []byte
	Expect []byte
}{
	{[]byte{'h', 'e', 'l', 'l', 'o', 0xC2, 0x85}, []byte{'h', 'e', 'l', 'l', 'o', ' '}},
	{[]byte{'h', 'o', 't', 0xC2, 0x85, 's', 't', 'u', 'f', 'f'}, []byte{'h', 'o', 't', ' ', 's', 't', 'u', 'f', 'f'}},
}

func TestRotate(t *testing.T) {
	for _, testCase := range testCases {
		compWS(&testCase.Input)
		if !reflect.DeepEqual(testCase.Input, testCase.Expect) {
			t.Errorf("invalid result. testCaes:%v", testCase)
		}
	}
}
