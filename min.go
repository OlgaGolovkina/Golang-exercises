package main

import "fmt"

func main()  {
	min :=[]int{2, 6, 8, 1}
	fmt.Println(MinNumb(min))

	max := []int{5, 25, 66, 6}
	fmt.Println(MaxNumb(max))
}

func MinNumb(xs []int) (min int) {
	min = xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return
}

// the same one decision for the max function

func MaxNumb(xs []int) int {
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}