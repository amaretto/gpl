package main

import "fmt"

func main() {
	num := 10
	fmt.Println("echo ", num, " :", echo(num))
}

func echo(num int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = num
		}
	}()
	panic("hoge")
}
