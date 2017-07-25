package main

import "fmt"

type Square struct {
	Side int
}

func(s *Square) Area() (area int)  {
	area = s.Side * s.Side
	return
}

func main()  {
	side := Square{5}
	fmt.Println(side)
	fmt.Println(side.Area())
}