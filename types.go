package main

import "fmt"

var a string

type b int

var c  b

func main()  {
	fmt.Printf("The type of a is %T\n", a)
	fmt.Printf("The type of b is %T\n", b(1))
	fmt.Printf("The type of a is %T\n", c)
}
