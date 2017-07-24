package main

import "fmt"

func main()  {
	monthdays := map[string]int {
		"January" : 31, "February" : 28, "March" : 31,
		"April" : 30, "May" : 31, "June" : 30,
		"July" : 31, "August" : 31, "September" : 30,
		"October" : 31, "November" : 30, "December" : 31,
	}

	year := 0
	for _, v := range monthdays {
		year += v
	}
	fmt.Printf("Number of days in the year is: %d\n",year)
}

