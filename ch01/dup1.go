// if you use Ctrl+D for interrupting input, result of this code may contains "D"

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Println(counts[input.Text()])
	}
	// Warning : this code doesn't care input.Error() occuring
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
