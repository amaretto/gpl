package main

import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s - s[slash_1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
