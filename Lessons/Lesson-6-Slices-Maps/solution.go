package main

import (
	"fmt"
)

func printSlice(slice []string) {
	fmt.Println("Slice Contents:")
	for i, item := range slice {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func printMap(data map[string]string) {
	fmt.Println("Map Contents:")

	for key, value := range data {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {
	// Task 1: Slice Operations
	numbers := []int{1, 2, 3, 4, 5}

	// Append Values
	numbers = append(numbers, 6, 7, 8)

	// Remove "3" (index 2)
	fmt.Println(numbers[:3], numbers[4:])
	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println("Final Slice:", numbers)

	// Task 2: Map Operations
	capitals := map[string]string{
		"USA":    "Washington DC",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	// Add India
	capitals["India"] = "New Delhi"

	// Delete France
	delete(capitals, "France")

	if _, exists := capitals["Germany"]; exists {
		fmt.Println("Germany is in the map.")
	} else {
		fmt.Println("Germany is not on the map.")
	}

	fmt.Println("Final Map:", capitals)

	chatters := []string{"vano2903", "docolero", "CrayonCrayoff", "Captnadmn", "DeepSquidGaming"}
	printSlice(chatters)
	printMap(capitals)
}
