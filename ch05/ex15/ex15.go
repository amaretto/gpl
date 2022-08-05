package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(10, 500, 20, 30))
	fmt.Println(max(10, 500, 20, 30))
}

func max(vals ...int) int {
	maxVal := -9223372036854775808
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func maxNeedArgs(vals ...int) (int, error) {
	maxVal := -9223372036854775808
	if len(vals) == 0 {
		return maxVal, fmt.Errorf("no args specified")
	}
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

func min(vals ...int) int {
	minVal := 9223372036854775807
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func minNeedArgs(vals ...int) (int, error) {
	minVal := 9223372036854775807
	if len(vals) == 0 {
		return minVal, fmt.Errorf("no args specified")
	}
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal, nil
}
