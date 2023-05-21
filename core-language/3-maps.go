package main

import "fmt"

func main() {
	// users := map[string]int{} // empty map
	users := map[string]int{
		"pascal": 23,
		"jude":   45,
		"mamie":  65,
	} // initialize a map

	employees := make(map[string]int) // empty map
	employees["Pascal"] = 23          // assign element to the map

	age, ok := employees["Pascal"]

	if !ok {
		fmt.Println("pascal not in the list")
	}

	fmt.Printf("users: %+v\n", users)
	fmt.Printf("employees: %+v\n", employees)
	fmt.Printf("employees age: %+v\n", age)

	delete(users, "pascal")
	fmt.Printf("users: %+v\n", users)

	for k, v := range users {
		fmt.Printf("My name is %s. I am %d years old.\n", k, v)
	}
}
