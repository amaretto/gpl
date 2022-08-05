package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/amaretto/study-go/golang_study2/ch02/lengthconv"
	"github.com/amaretto/study-go/golang_study2/ch02/tempconv"
	"github.com/amaretto/study-go/golang_study2/ch02/weightconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 100000000
)

func main() {
	var args []string
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	} else {
		args = os.Args[1:]
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		// templature
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		// weight
		k := weightconv.Kilogram(t)
		p := weightconv.Pond(t)
		fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
		// length
		m := lengthconv.Metre(t)
		feet := lengthconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToF(m), feet, lengthconv.FToM(feet))
	}
}
