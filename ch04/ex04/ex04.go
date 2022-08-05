package main

const Length = 10

func rotate(s []int, p int) []int {
	result := make([]int, len(s))
	copy(result, s[p:])
	copy(result[len(s)-p:], s[:p])
	return result
}
