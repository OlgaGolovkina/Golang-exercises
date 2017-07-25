package main

import "fmt"

func average(xs []float64) (avg float64)  {
	sum := 0.0

	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum +=v
		}
		avg = sum/float64(len(xs))
	}
	return
}

func main()  {
	xs := []float64{2.3, 1.2, 66.2}
	fmt.Println(average(xs))
}