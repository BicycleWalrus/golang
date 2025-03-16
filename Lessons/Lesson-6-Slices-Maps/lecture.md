# **Lesson 6: Slices and Maps in Go**

### **Why This Lesson?**
Now that we’ve covered control flow, it’s time to explore **data structures** in Go. This lesson will introduce:
✅ **Slices** – Go’s dynamic arrays  
✅ **Maps** – Key-value pairs (like Python dictionaries)  
✅ How to manipulate slices and maps

---

## **Lesson Outline**
### **1. Slices in Go**
- What are slices?
- Declaring and initializing slices
- Adding elements (`append`)
- Removing elements (workaround)
- Iterating over slices

### **2. Maps in Go**
- What are maps?
- Declaring and initializing maps
- Adding, accessing, and deleting values
- Checking if a key exists
- Iterating over a map

---

## **1. Slices in Go**
### **What is a Slice?**
- A **slice** is a dynamically sized array.
- Unlike arrays, slices do **not** have a fixed length.
- More similar to Python **lists** than arrays.

### **Declaring and Initializing Slices**
```go
package main

import "fmt"

func main() {
    // Declaring a slice (empty)
    var numbers []int

    // Initializing a slice with values
    fruits := []string{"Apple", "Banana", "Cherry"}

    fmt.Println("Numbers:", numbers)  // Empty slice
    fmt.Println("Fruits:", fruits)    // ["Apple", "Banana", "Cherry"]
}
```

### **Adding Elements to a Slice (`append`)**
```go
func main() {
    var numbers []int // Empty slice

    numbers = append(numbers, 10)  // Add 10
    numbers = append(numbers, 20, 30)  // Add multiple values

    fmt.Println(numbers)  // Output: [10 20 30]
}
```

### **Removing Elements from a Slice (Workaround)**
Go does **not** have a built-in `remove` function.  
To remove an element at index `i`, you can **slice around it**:
```go
func main() {
    fruits := []string{"Apple", "Banana", "Cherry", "Date"}

    // Remove "Banana" (index 1)
    fruits = append(fruits[:1], fruits[2:]...)

    fmt.Println(fruits)  // Output: ["Apple", "Cherry", "Date"]
}
```

### **Iterating Over a Slice**
```go
func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}

    for index, fruit := range fruits {
        fmt.Println(index, fruit)
    }
}
```
> Use `_` to **omit** the index:
```go
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

---

## **2. Maps in Go**
### **What is a Map?**
- A **map** is a key-value store (like a Python dictionary).
- Keys must be of a comparable type (strings, numbers, etc.).
- Values can be any type.

### **Declaring and Initializing a Map**
```go
func main() {
    // Declare and initialize a map
    person := map[string]int{
        "Alice": 25,
        "Bob":   30,
    }

    fmt.Println(person)  // Output: map[Alice:25 Bob:30]
}
```

### **Adding and Accessing Map Values**
```go
func main() {
    ages := make(map[string]int)  // Create an empty map

    // Adding key-value pairs
    ages["Alice"] = 25
    ages["Bob"] = 30

    fmt.Println(ages["Alice"])  // Output: 25
}
```

### **Checking If a Key Exists**
```go
func main() {
    person := map[string]int{"Alice": 25}

    age, exists := person["Bob"]  // Bob is not in the map
    if exists {
        fmt.Println("Bob's age is", age)
    } else {
        fmt.Println("Bob is not found")
    }
}
```

### **Deleting a Key from a Map**
```go
func main() {
    person := map[string]int{"Alice": 25, "Bob": 30}

    delete(person, "Bob")  // Remove "Bob"
    fmt.Println(person)  // Output: map[Alice:25]
}
```

### **Iterating Over a Map**
```go
func main() {
    person := map[string]int{"Alice": 25, "Bob": 30}

    for name, age := range person {
        fmt.Println(name, "is", age, "years old.")
    }
}
```

---

## **3. Hands-On Exercise**
### 📝 **Task 1: Slice Manipulation**
1. Create a slice `numbers := []int{1, 2, 3, 4, 5}`
2. Append `6, 7, 8` to it.
3. Remove the number `3` from the slice.
4. Print the final slice.

### 📝 **Task 2: Map Operations**
1. Create a map `capitals := map[string]string{"USA": "Washington", "France": "Paris"}`
2. Add `"Japan": "Tokyo"`
3. Delete `"France"`
4. Check if `"India"` exists in the map.
5. Print the final map.

---

## **Conclusion**
- **Slices** are dynamic lists that can grow/shrink.
- **Maps** store key-value pairs, like dictionaries.
- **Manipulation** involves `append`, slicing, and iteration.