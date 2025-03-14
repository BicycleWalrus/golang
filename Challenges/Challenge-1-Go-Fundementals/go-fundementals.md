# **Review Exercise: Go Fundamentals**

## **Objective**
This review exercise will reinforce the key concepts you've learned so far:  
✅ Printing and formatting output  
✅ Working with variables and constants  
✅ Understanding data types  
✅ Using functions and return values  
✅ Handling multiple return values  
✅ Customizing print output

---

## **Exercise Instructions**
### 📝 **Task 1: Printing with Custom Separators**
1. Write a function `customPrint` that takes a **separator** (string) and a list of **words** (slice of strings).
2. The function should print all words joined by the separator.
3. Example usage:
   ```go
   customPrint("-", "Go", "is", "fun") 
   ```
   **Expected output:**
   ```
   Go-is-fun
   ```

---

### 📝 **Task 2: Variable Declarations**
1. Declare and initialize the following variables:
    - An **integer** named `year` with the value `2025`.
    - A **constant** named `pi` with the value `3.14159`.
    - A **boolean** named `isGoFun` set to `true`.
2. Print them all on one line using a **comma** as a separator.

---

### 📝 **Task 3: Functions & Return Values**
1. Write a function `addNumbers` that takes **two integers** as input and returns their **sum**.
2. Write another function `swapValues` that takes **two strings** and returns them **swapped**.
3. Call both functions in `main()` and print the results.

---

### 📝 **Task 4: Multiple Return Values**
1. Create a function `mathOperations` that takes two numbers and returns:
    - Their sum
    - Their product
    - Their difference
    - Their quotient
2. Call this function in `main()` and print the results clearly.

---

## **Solution Template (Fill in the Blanks)**

```go
package main

import (
    "fmt"
    "strings"
)

// Task 1: Custom Print Function
func customPrint(separator string, words ...string) {
    fmt.Println(strings.Join(words, separator))
}

// Task 3: Function to Add Two Numbers
func addNumbers(a int, b int) int {
    return // Your code here
}

// Task 3: Function to Swap Two Strings
func swapValues(x string, y string) (string, string) {
    return // Your code here
}

// Task 4: Function with Multiple Return Values
func mathOperations(a int, b int) (int, int, int, float64) {
    sum := // Your code here
    product := // Your code here
    difference := // Your code here
    quotient := // Your code here
    return sum, product, difference, quotient
}

func main() {
    // Task 2: Variables
    var year =  // Your code here
    const pi =  // Your code here
    isGoFun :=  // Your code here

    fmt.Println(year, pi, isGoFun)  // Print variables using commas

    // Task 1: Call Custom Print Function
    customPrint("-", "Go", "is", "fun")

    // Task 3: Call Functions and Print Results
    sum := addNumbers(10, 5)
    fmt.Println("Sum:", sum)

    first, second := swapValues("Hello", "World")
    fmt.Println("Swapped:", first, second)

    // Task 4: Call Function and Print Multiple Return Values
    s, p, d, q := mathOperations(10, 3)
    fmt.Println("Sum:", s, "Product:", p, "Difference:", d, "Quotient:", q)
}
```

---

## **Expected Output (Example)**
```
2025 3.14159 true
Go-is-fun
Sum: 15
Swapped: World Hello
Sum: 13 Product: 30 Difference: 7 Quotient: 3.333333
```

---
### 🚀 **Challenge Extension:**
If this was too easy, try modifying:
1. **Add Error Handling** – What happens if `mathOperations` is given `b = 0`?
2. **Extend `customPrint`** – Allow it to handle `int` and `float` types too.
3. **Use `fmt.Sprintf`** – Format your output with custom labels.

---
Would you like a **graded quiz format** next time, or do you prefer these hands-on exercises? 😃