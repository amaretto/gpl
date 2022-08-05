package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	categories := make(map[string]int)

	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// from graphic.go const
		// pC     = 1 << iota // a control character.
		// pP                 // a punctuation character.
		// pN                 // a numeral.
		// pS                 // a symbolic character.
		// pZ                 // a spacing character.
		// pLu                // an upper-case letter.
		// pLl                // a lower-case letter.
		// pp                 // a printable character according to Go's definition.
		// pg     = pp | pZ   // a graphical character according to the Unicode definition.
		// pLo    = pLl | pLu // a letter that is neither upper nor lower case.
		if unicode.IsControl(r) {
			categories["control"]++
		} else if unicode.IsPunct(r) {
			categories["punctuaion"]++
		} else if unicode.IsNumber(r) {
			categories["number"]++
		} else if unicode.IsSymbol(r) {
			categories["symbol"]++
		} else if unicode.IsSpace(r) {
			categories["space"]++
		} else if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				categories["upper letter"]++
			} else if unicode.IsLower(r) {
				categories["lower letter"]++
			} else {
				categories["letter"]++
			}
		} else if unicode.IsPrint(r) {
			categories["printable"]++
		} else if unicode.IsDigit(r) {
			categories["digit"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\ncategory\tcount\n")
	for c, n := range categories {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 character\n", invalid)
	}
}
