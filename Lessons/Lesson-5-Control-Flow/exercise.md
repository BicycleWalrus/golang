# **Lesson 5: Control Flow – Exercise**

## **Objective**
This exercise will test your understanding of:
✅ Conditional statements (`if`, `else if`, `else`)  
✅ `switch` statements  
✅ Loops (`for`, `range`)  
✅ Loop control (`break`, `continue`)

---

## **Instructions**
### 📝 **Task 1: Age Categorizer (if-else)**
1. Write a function `categorizeAge(age int) string` that:
    - Returns `"Child"` if `age` is less than **13**.
    - Returns `"Teenager"` if `age` is between **13 and 19**.
    - Returns `"Adult"` otherwise.
2. Call this function in `main()` with different ages and print the results.

---

### 📝 **Task 2: Number Classification (switch statement)**
1. Write a function `classifyNumber(num int) string` that:
    - Uses a **switch statement** to classify numbers:
        - `"Even"` if `num` is divisible by 2.
        - `"Odd"` otherwise.
2. Call this function in `main()` with different numbers and print the results.

---

### 📝 **Task 3: While-like Loop (for loop)**
1. Write a `for` loop that simulates a **while loop** to print numbers **1 to 10**.
2. If the number is a **multiple of 3**, **skip it** using `continue`.
3. If the number is **greater than 8**, **stop the loop** using `break`.

---

### 📝 **Task 4: Loop Through a Slice**
1. Create a slice of **colors**: `["Red", "Green", "Blue", "Yellow"]`
2. Loop through the slice using `range`, but **omit the index** when printing.

---

## **Solution Template**
```go
package main

import (
    "fmt"
)

// Task 1: Age Categorizer
func categorizeAge(age int) string {
    // Your code here
}

// Task 2: Number Classification
func classifyNumber(num int) string {
    // Your code here
}

func main() {
    // Task 1: Test Age Categorizer
    fmt.Println(categorizeAge(10))  // Expected: Child
    fmt.Println(categorizeAge(16))  // Expected: Teenager
    fmt.Println(categorizeAge(25))  // Expected: Adult

    // Task 2: Test Number Classification
    fmt.Println(classifyNumber(8))  // Expected: Even
    fmt.Println(classifyNumber(7))  // Expected: Odd

    // Task 3: While-like loop
    num := 1
    for {
        if num > 10 {
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
```

---

## **Expected Output**
```
Child
Teenager
Adult
Even
Odd
1
2
4
5
7
8
Red
Green
Blue
Yellow
```

---

### 🚀 **Bonus Challenge**
- Modify `categorizeAge()` to handle **edge cases**, like negative ages.
- Enhance `classifyNumber()` to **handle zero separately**.
- Add **user input** using `fmt.Scanln()` to test your functions interactively.

Would you like a **full solution** provided now, or do you want to attempt it first? 😃