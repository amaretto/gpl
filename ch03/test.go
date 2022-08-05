package main

import "fmt"

func main() {
	s := "hello world!!"
	s2 := s[6:]

	fmt.Printf("s:%s\n", s)
	fmt.Printf("s2:%s\n", s2)

	s = "world hello"

	fmt.Printf("s:%s\n", s)
	fmt.Printf("s2:%s\n", s2)
}
