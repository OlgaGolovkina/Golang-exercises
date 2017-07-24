package main

import "fmt"

type MyObj struct {
	Name    string
	Surname string
	age     int64
}

func main()  {
	myObject := MyObj{"Dave","Tompson", 42}
	fmt.Println(myObject)
	fmt.Println(myObject.ChangeName())
}

func(m *MyObj)ChangeName() string{
	 m.Name = "Kate"
	return m.Name
}