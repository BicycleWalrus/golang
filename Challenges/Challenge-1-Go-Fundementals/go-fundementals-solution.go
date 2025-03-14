package main

import (
	"fmt"
	"strings"
)

func customPrint(separator string, words ...string) {
	fmt.Println(strings.Join(words, separator))
}

func addNumbers(x int, y int) int {
	return x + y
}

func swapValues(a int, b int) (int, int) {
	return b, a
}

func mathOperations(x int, y int) (int, int, int, float64) {
	sum := x + y
	product := x * y
	diffenece := x - y
	var quotient float64

	if y != 0 { // Prevent division by zero
		quotient = float64(x) / float64(y)
	} else {
		fmt.Println("Warning: Division by zero is not allowed.")
		quotient = 0.0
	}
	return sum, product, diffenece, quotient
}

func main() {
	// Task 2: Variables
	var year = 2025
	const pi = 3.14159
	var isGoFun bool = true

	customPrint(", ", fmt.Sprint(year), fmt.Sprint(pi), fmt.Sprint(isGoFun))

	// Task 1: Custom Print Function
	customPrint("-", "Go", "Is", "fun")

	// Task 3: Functions & Return Values
	fmt.Println(addNumbers(21, 9))
	fmt.Println(swapValues(21, 11))

	// Task 4: Multiple Return Values
	s, p, d, q := mathOperations(21, 9)
	fmt.Println("Sum:", s, "Product:", p, "Difference:", d, "Quotient:", q)

}
