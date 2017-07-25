package main

import "fmt"

type calculator interface {
	add() int64
	sub() float64
	mul() int64
	div() float64
}


type values struct {
	x, y int
}

func (a values) add() int {
	return a.x + a.y
}

func (a values) sub() int {
	return a.x - a.y
}
func (a values) mul() int {
	return a.x * a.y
}
func (a values) div() int {
	return a.x / a.y
}

func main() {

	a := values{x: 30, y: 20}

	fmt.Println("values: ", a)
	fmt.Println("addition: ", a.add())
	fmt.Println("subtration: ", a.sub())
	fmt.Println("multiplication: ", a.mul())
	fmt.Println("division: ", a.div())
}
