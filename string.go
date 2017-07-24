package main

import "fmt"

func main() {
	str1 := "palm"
	character := []rune(str1)
	character[3] = 'e'

	str2 := string(character)
	fmt.Printf("Changing word '%s' to '%s'\n" , str1, str2)
}
