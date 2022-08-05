package main

import (
	"strings"
)

func useJoin(input []string) string {
	return strings.Join(input, " ")
}

func useOperand(input []string) string {
	var s, sep string
	for i := 0; i < len(input); i++ {
		s += sep + input[i]
		sep = " "
	}
	return s
}
