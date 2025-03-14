package main

import "fmt"

func greet() {
	fmt.Println("Hello from Go!")
}

func add(x int, y int) int {
	return x + y
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func rectangleDimensions(length int, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Explicit return with named values
}

func main() {
	greet()

	sum := add(7, 3)
	fmt.Println("Sum:", sum)

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	a, p := rectangleDimensions(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)
}
