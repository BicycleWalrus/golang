package main

import (
	"fmt"
)

func gradeCalc(score int) string {
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 && score < 90 {
		grade = "B"
	} else if score >= 70 && score < 80 {
		grade = "C"
	} else if score >= 60 && score < 70 {
		grade = "D"
	} else {
		grade = "E"
	}
	return grade
}

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You're an Adult")
	} else {
		fmt.Println("You're a Child.")
	}

	temperature := 75

	if temperature > 80 {
		fmt.Println("It's hot!")
	} else {
		fmt.Println("It's not hot!")
	}

	fmt.Println("Grade:", gradeCalc(85))
}
