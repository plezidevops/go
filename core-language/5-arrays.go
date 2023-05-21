package main

import "fmt"

func main() {
	numbers := [3]int{} // emty slice

	numbers[0] = 3
	numbers[1] = 4
	numbers[2] = 9

	fmt.Printf("The numbers %+v\n", numbers)

	for v := range numbers {
		fmt.Printf("This number is: %d\n", v)
	}

}
