package main

import "fmt"

func main() {
	fmt.Println(max_number(1, 2, 3, 4, 5))
}

func max_number(args ...int) int {
	var maxNumber int = 0
	for _, number := range args {
		if number > maxNumber {
			maxNumber = number
		}
	}
	return maxNumber
}