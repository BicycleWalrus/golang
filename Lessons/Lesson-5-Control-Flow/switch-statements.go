package main

import "fmt"

func eoro(x int) string {
	num := x
	var eoro string

	switch {
	case num%2 == 0:
		eoro = "Even Number"
	case num%2 != 0:
		eoro = "Odd Number"
	}
	return eoro
}

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("You've got a case of the Mondays.")
	case "Tuesday":
		fmt.Println("It's Taco Tuesday")
	case "Wednesday":
		fmt.Println("It's Hump Day!")
	default:
		fmt.Println("It's a regular day.")
	}

	fmt.Println(eoro(3))
}
