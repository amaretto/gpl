package main

import (
	"flag"
	"fmt"

	"github.com/amaretto/study-go/golang_study2/ch07/ex06/tempflag"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temprature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
