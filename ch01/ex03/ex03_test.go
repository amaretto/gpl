package main

import "testing"

var testCases = []struct {
	Input  []string
	Output string
}{
	{[]string{"Raising", "Elephant", "Is", "So", "Utterly", "Boring"}, "Raising Elephant Is So Utterly Boring"},
	{[]string{"1", "2", "3", "4"}, "1 2 3 4"},
}

func TestUseJoin(t *testing.T) {
	for _, testCase := range testCases {
		result := useJoin(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestUseOperand(t *testing.T) {
	for _, testCase := range testCases {
		result := useOperand(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func BenchmarkUseJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useJoin(testCases[0].Input)
	}
}

func BenchmarkUseOperand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useOperand(testCases[0].Input)
	}
}
