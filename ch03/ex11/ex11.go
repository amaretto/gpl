package main

import (
	"bytes"
)

func commaWithBB(s string) string {
	b := []byte(s)
	var bdecimal []byte
	var bb bytes.Buffer

	if bytes.HasPrefix(b, []byte{'-'}) {
		bb.Write([]byte{'-'})
		b = b[1:]
	}
	if dot := bytes.LastIndex(b, []byte{'.'}); dot >= 0 {
		bdecimal = b[dot:]
		b = b[:dot]
	}

	for i := 0; i < len(b); i++ {
		if i != 0 && (len(b)-i)%3 == 0 {
			bb.Write([]byte{','})
		}
		bb.Write([]byte{b[i]})
	}
	bb.Write(bdecimal)
	return bb.String()

}
