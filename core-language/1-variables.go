package main

import "fmt"

// global variables
var (
	firstName = "Jude"
	lastName  = "Pierre"
)

// constant
const version = 1

func main() {

	fmt.Println(firstName, lastName)

	// can't modified a constant variable
	version = 10

}
