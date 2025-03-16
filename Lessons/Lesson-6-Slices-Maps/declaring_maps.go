package main

import "fmt"

func main() {
	person := map[string]int{
		"James": 32,
		"Alice": 34,
		"Bob":   23,
	}
	fmt.Println(person)
	fmt.Println(person["James"])

	ages := make(map[string]int) // Make an Empty Map

	fmt.Println(ages)
	ages["James"] = 32
	fmt.Println(ages)
	ages["Alice"] = 34
	fmt.Println(ages)
	ages["Bob"] = 23
	fmt.Println(ages)

	fmt.Println(ages["Bob"]) // Expected Value: 23

	age, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age is", age)
	} else {
		fmt.Println("Bob doesn't exist.")
	}

	age, exists = ages["Charlie"]
	if exists {
		fmt.Println("Charlie's age is", age)
	} else {
		fmt.Println("Charlie doesn't exist.")
	}

	delete(ages, "Bob") // Exterminate Bob
	fmt.Println(ages)
}
