package main

import (
	"fmt"
	"os"
)

func main() {
	for n, arg := range os.Args[1:] {
		fmt.Printf("%d, %s\n", n, arg)
	}
}
