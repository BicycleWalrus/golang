# **Lesson 6: Slices and Maps – Exercise**

## **Objective**
This exercise will reinforce your understanding of:
✅ Slices – Creating, modifying, and iterating  
✅ Maps – Storing, accessing, and removing key-value pairs

---

## **Exercise Instructions**

### 📝 **Task 1: Slice Operations**
1. Create a slice called `numbers` with values **1, 2, 3, 4, 5**.
2. Append the values **6, 7, and 8** to the slice.
3. Remove the number **3** from the slice without using a built-in `remove` function.
4. Print the final slice.

---

### 📝 **Task 2: Map Operations**
1. Create a map called `capitals` to store **country-capital** pairs:
    - `"USA": "Washington"`
    - `"France": "Paris"`
    - `"Japan": "Tokyo"`
2. Add a new key-value pair `"India": "New Delhi"`.
3. Delete the `"France"` entry.
4. Check if `"Germany"` exists in the map and print a message accordingly.
5. Print the final map.

---

### 📝 **Task 3: Iterating Over Data**
1. Write a function `printSlice(slice []string)` that:
    - Accepts a slice of strings as input.
    - Prints each element on a new line.
2. Write a function `printMap(data map[string]string)` that:
    - Accepts a map as input.
    - Prints `"The capital of <Country> is <Capital>"` for each key-value pair.

---

## **Solution Template (Fill in the Blanks)**
```go
package main

import (
    "fmt"
)

// Task 3: Function to Print a Slice
func printSlice(slice []string) {
    // Your code here
}

// Task 3: Function to Print a Map
func printMap(data map[string]string) {
    // Your code here
}

func main() {
    // Task 1: Slice Operations
    numbers := []int{1, 2, 3, 4, 5}

    // Append values
    numbers = append(numbers, 6, 7, 8)

    // Remove "3" (index 2)
    numbers = append(numbers[:2], numbers[3:]...)

    fmt.Println("Final Slice:", numbers)

    // Task 2: Map Operations
    capitals := map[string]string{
        "USA":    "Washington",
        "France": "Paris",
        "Japan":  "Tokyo",
    }

    // Add new entry
    capitals["India"] = "New Delhi"

    // Delete "France"
    delete(capitals, "France")

    // Check if "Germany" exists
    if _, exists := capitals["Germany"]; exists {
        fmt.Println("Germany is in the map.")
    } else {
        fmt.Println("Germany is not found in the map.")
    }

    // Print final map
    fmt.Println("Final Map:", capitals)

    // Task 3: Call print functions
    colors := []string{"Red", "Green", "Blue"}
    printSlice(colors)

    printMap(capitals)
}
```

---

## **Expected Output**
```
Final Slice: [1 2 4 5 6 7 8]
Germany is not found in the map.
Final Map: map[India:New Delhi Japan:Tokyo USA:Washington]
Red
Green
Blue
The capital of USA is Washington.
The capital of Japan is Tokyo.
The capital of India is New Delhi.
```

---

### 🚀 **Bonus Challenge**
- Modify `printSlice` to print a **numbered list** (e.g., `1. Red`).
- Modify `printMap` to **sort keys alphabetically** before printing.
- Add user input for **country lookups**.