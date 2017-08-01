package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	Name string
	Age int
}

func main() {
	u1 := user{
		Name: "Jake",
		Age: 32,
	}

	u2 := user{
		Name: "Sonya",
		Age: 45,
	}

	u3 := user{
		Name: "Kate",
		Age: 25,
	}

	users := []user{u1, u2, u3,}
	fmt.Println(users)

// marshal the []user to JSON
	byteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(byteSlice)

	jsonUsers := `[{"Name":"Jake","Age":32},{"Name":"Sonya","Age":45},{"Name":"Kate","Age":25}]`
	fmt.Println(jsonUsers)
}