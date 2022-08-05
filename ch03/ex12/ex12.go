package main

import (
	"sort"
)

type sortRune []rune

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRune(r))
	return string(r)
}

func anagram(s1, s2 string) bool {
	ss1 := SortString(s1)
	ss2 := SortString(s2)
	if ss1 == ss2 {
		return true
	} else {
		return false
	}
}
