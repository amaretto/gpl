package main

import (
	"unicode"
)

func compWS(str *[]byte) {
	var s = *str

	for i := 0; i < len(s); i++ {
		if 192 <= s[i] && s[i] < 224 {
			if unicode.IsSpace([]rune(string(s[i : i+2]))[0]) {
				s[i] = byte(' ')
				copy(s[i+1:], s[i+2:])
				s = s[:len(s)-1]
			}
		}
	}
	*str = s
}
