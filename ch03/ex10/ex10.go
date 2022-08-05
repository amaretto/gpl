package main

import (
	"bytes"
)

func commaWithBB(s string) string {
	var bb bytes.Buffer

	for i := 0; i < len(s); i++ {
		if i != 0 && (len(s)-i)%3 == 0 {
			bb.Write([]byte{','})
		}
		bb.Write([]byte{s[i]})
	}
	return bb.String()

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
