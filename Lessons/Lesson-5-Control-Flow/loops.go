package main

import "fmt"

func main() {
	// Basic For Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Looping over a Slice
	fruits := []string{"Apple", "Banana", "Cherry"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Infinite Loop
	counter := 0

	for {
		counter++
		fmt.Println("Looping...", counter)
		if counter == 3 {
			break
		}
	}

	// GoLang While Loop
	number := 1
	for number <= 5 {
		fmt.Println("Golang:", number)
		number++
	}

	// Loop Continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Control:", i)
	}
}
