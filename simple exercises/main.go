package main

import "fmt"

func main()  {
	var input string
	fmt.Println("What's your name?")
	fmt.Scanln(&input)
	fmt.Printf("Hello, %s", input)
}