package main

import "fmt"

func main() {
	numbers := []int{} // emty slice
	otherNumbers := []int{1, 2, 3}
	otherNumbers2 := make([]int, 0)

	numbers = append(numbers, 3)

	fmt.Printf("The numbers %+v\n", numbers)
	fmt.Printf("The numbers %+v\n", otherNumbers)
	fmt.Printf("The numbers %+v\n", otherNumbers2)
}
