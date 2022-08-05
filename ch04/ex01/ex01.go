package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCountSHA256(c1, c2 [32]byte) int {
	var c int
	for i := 0; i < 32; i++ {
		c += int(pc[c1[i]^c2[i]])
	}
	return c
}

func main() {
	var c1, c2 [32]byte
	c1[0] = byte(5)
	c2[0] = byte(3)
	fmt.Println(popCountSHA256(c1, c2))
}
