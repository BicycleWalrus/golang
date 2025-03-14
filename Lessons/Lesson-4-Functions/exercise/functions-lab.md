## Functions Exercise
1. Write a function `multiply` that takes two integers and returns their product.
2. Create a function `swap` that swaps two string values.
3. Call these functions in `main()` and print the results.

**Solution:**
```go
package main

import "fmt"

// Function to multiply two numbers
func multiply(a int, b int) int {
    return a * b
}

// Function to swap two strings
func swap(x string, y string) (string, string) {
    return y, x
}

func main() {
    product := multiply(4, 5)
    fmt.Println("Product:", product)

    first, second := swap("hello", "world")
    fmt.Println("Swapped:", first, second)
}
```