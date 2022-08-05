package main

import (
	"sort"
)

type byRune []rune

func (x byRune) Len() int           { return len(x) }
func (x byRune) Less(i, j int) bool { return x[i] < x[j] }
func (x byRune) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func isPalindrome(s sort.Interface) bool {
	var j int

	for i := 0; i < s.Len()/2; i++ {
		j = s.Len() - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}
