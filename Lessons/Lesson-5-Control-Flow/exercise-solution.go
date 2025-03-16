package main

import (
	"fmt"
)

// Task 1: Age Categorizer
func categorizeAge(age int) string {
	if age < 13 {
		return "You're a Child"
	} else if age >= 13 && age < 18 {
		return "You're a Teenager!"
	} else {
		return "You're an Adult!"
	}
}

// Task 2: Number Classification
func classifyNumber(num int) string {
	var eoro string

	switch {
	case num == 0:
		eoro = "Zero"
	case num%2 == 0:
		eoro = "Even Number"
	case num%2 != 0:
		eoro = "Odd Number"
	}
	return eoro
}

func main() {
	// Task 1: Test Age Categorizer
	fmt.Println(categorizeAge(10)) // Expected: Child
	fmt.Println(categorizeAge(16)) // Expected: Teenager
	fmt.Println(categorizeAge(25)) // Expected: Adult

	// Task 2: Test Number Classification
	fmt.Println(classifyNumber(8)) // Expected: Even
	fmt.Println(classifyNumber(7)) // Expected: Odd

	// Task 3: While-like loop
	num := 1
	for num <= 10 {
		if num > 8 {
			break
		}
		if num%3 == 0 {
			num++
			continue
		}
		fmt.Println(num)
		num++
	}

	// Task 4: Loop through a slice
	colors := []string{"Red", "Green", "Blue", "Yellow"}
	for _, color := range colors {
		fmt.Println(color)
	}
}
