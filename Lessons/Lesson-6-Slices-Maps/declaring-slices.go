package main

import "fmt"

func main() {
	// Declaring a Slice (empty)
	var numbers []int
	fmt.Println("Numbers:", numbers)

	// Add elements to the slice
	numbers = append(numbers, 10)
	fmt.Println("Numbers:", numbers)
	numbers = append(numbers, 20, 30, 40)
	fmt.Println("Numbers:", numbers)

	// Initializing a slice with values
	fruits := []string{"peach", "banana", "kiwi", "date"}
	fmt.Println("Fruits:", fruits)
	fruits = append(fruits[:1], fruits[2:]...) // Remove Banana, or Index 1
	fmt.Println("Fruits:", fruits)

	// Iterating over a Slice
	for _, fruit := range fruits {
		fmt.Println("Fruit:", fruit)
	}
}
