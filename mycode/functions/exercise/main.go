package main

import (
	"fmt"
)

func multiply(x int, y int) int {
	return x * y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	product := multiply(4, 17)
	fmt.Println(product)

	first, second := swap("Bicycle", "Walrus")
	fmt.Print("Name: ", first, ", ", second)
}
