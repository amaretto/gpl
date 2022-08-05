package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press return to abort.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(1 * time.Second)
	// ここがダメ
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-ticker.C
	}
	select {
	case <-abort:
		fmt.Println("Launch aborted!")
		ticker.Stop()
		return
	default:
		// nothing to do
	}

	launch()
}

func launch() {
	fmt.Println("launch")
}
