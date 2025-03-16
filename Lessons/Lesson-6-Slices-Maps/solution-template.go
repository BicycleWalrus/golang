package main

import (
	"fmt"
)

// Task 3: Function to Print a Slice
func printSlice(slice []string) {
	// Your code here
}

// Task 3: Function to Print a Map
func printMap(data map[string]string) {
	// Your code here
}

func main() {
	// Task 1: Slice Operations
	numbers := []int{1, 2, 3, 4, 5}

	// Append values
	numbers = append(numbers, 6, 7, 8)

	// Remove "3" (index 2)
	numbers = append(numbers[:2], numbers[3:]...)

	fmt.Println("Final Slice:", numbers)

	// Task 2: Map Operations
	capitals := map[string]string{
		"USA":    "Washington",
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	// Add new entry
	capitals["India"] = "New Delhi"

	// Delete "France"
	delete(capitals, "France")

	// Check if "Germany" exists
	if _, exists := capitals["Germany"]; exists {
		fmt.Println("Germany is in the map.")
	} else {
		fmt.Println("Germany is not found in the map.")
	}

	// Print final map
	fmt.Println("Final Map:", capitals)

	// Task 3: Call print functions
	colors := []string{"Red", "Green", "Blue"}
	printSlice(colors)

	printMap(capitals)
}
