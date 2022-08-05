package popcount

import "testing"

var testCases = []struct {
	Input  uint64
	Output int
}{
	{0, 0},
	{255, 8},
	{256, 1},
	{18446744073709551615, 64},
}

func TestPopCount(t *testing.T) {
	for _, testCase := range testCases {
		result := PopCount(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestPopCountUseLoop(t *testing.T) {
	for _, testCase := range testCases {
		result := PopCountUseLoop(testCase.Input)
		if testCase.Output != result {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testCases[0].Input)
	}
}

func BenchmarkPopCountUseLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountUseLoop(testCases[0].Input)
	}
}
