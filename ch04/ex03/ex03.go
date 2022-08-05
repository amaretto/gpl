package main

const Length = 10

func reverse(s *[Length]int) {
	for i, j := 0, Length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
