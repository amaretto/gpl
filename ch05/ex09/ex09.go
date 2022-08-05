package main

import (
	"fmt"
	"os"
	"strings"
)

var subs string

func main() {
	s := os.Args[1]
	result := expand(s, func(subs string) string {
		return strings.ToUpper(subs)
	})
	fmt.Println(result)
}

func expand(s string, f func(string) string) string {
	if s == "" {
		return ""
	}

	sarray := strings.Split(s, " ")
	for _, word := range sarray {
		if word[0] == '$' {
			target := string(word[1:])
			s = strings.Replace(s, word, f(target), 1)
		}
	}
	return s
}
