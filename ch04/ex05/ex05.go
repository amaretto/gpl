package main

import "fmt"

const Length = 10

func delDup(s []int) []int {
	var tmp, v int
	for i := 0; i < len(s); i++ {
		v = s[i]
		if i == 0 {
			tmp = v
			continue
		}
		if v == tmp {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
		tmp = v
	}
	return s
}

func main() {

	s := []int{0, 1, 2, 3, 3}

	fmt.Println(s)
	s = delDup(s)
	fmt.Println(s)
}
