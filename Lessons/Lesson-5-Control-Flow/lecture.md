## **Lesson 5: Control Flow in Go (Conditional Statements & Loops)**

### **Why This Lesson?**
Now that we've covered functions and variables, it's time to explore **control flow**—the foundation of decision-making in any programming language. This lesson will introduce:
✅ `if` / `else` statements  
✅ `switch` statements  
✅ `for` loops (Go’s only loop structure)  
✅ Loop control statements (`break`, `continue`)

---

## **Lesson Outline**

### **1. Conditional Statements (`if`, `else`, `else if`)**
Just like in Python, Go allows you to execute different code blocks based on conditions.

#### **Basic `if` Statement**
```go
package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("You are an adult.")
    }
}
```

#### **Using `if...else`**
```go
func main() {
    temperature := 25

    if temperature > 30 {
        fmt.Println("It's hot!")
    } else {
        fmt.Println("It's not too hot.")
    }
}
```

#### **Using `if...else if...else`**
```go
func main() {
    score := 85

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C or lower")
    }
}
```

---

### **2. `switch` Statements (Alternative to `if-else` Chains)**
`switch` is a cleaner way to handle multiple conditions.

#### **Basic `switch` Example**
```go
func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("Start of the week.")
    case "Friday":
        fmt.Println("Weekend is near!")
    default:
        fmt.Println("A regular day.")
    }
}
```

#### **Switch Without a Value (More Like `if` Statements)**
```go
func main() {
    num := 15

    switch {
    case num%2 == 0:
        fmt.Println("Even number")
    case num%2 != 0:
        fmt.Println("Odd number")
    }
}
```

---

### **3. Loops in Go (`for`)**
Unlike Python, Go has only **one** loop structure: `for`. There’s no `while` loop!

#### **Basic `for` Loop**
```go
func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }
}
```

#### **Looping Over a Slice (`range`)**
```go
func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}

    for index, fruit := range fruits {
        fmt.Println(index, fruit)
    }
}
```

#### **Infinite Loop (Similar to `while True:` in Python)**
```go
func main() {
    counter := 0

    for {
        fmt.Println("Looping...")
        counter++
        if counter == 3 {
            break
        }
    }
}
```

---

### **4. Loop Control Statements**
#### **Using `break`**
```go
func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            break // Exits the loop
        }
        fmt.Println(i)
    }
}
```
**Output:**
```
1
2
```

#### **Using `continue`**
```go
func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // Skips this iteration
        }
        fmt.Println(i)
    }
}
```
**Output:**
```
1
2
4
5
```

---

### **5. Hands-On Exercise**
#### **Task 1: Age Check**
Write a program that asks for an age and prints:
- `"Child"` if age < 13
- `"Teenager"` if age is between 13 and 19
- `"Adult"` otherwise

#### **Task 2: Loop Through Numbers**
Write a `for` loop that prints numbers **1 to 10**, but:
- Skips **multiples of 3** using `continue`
- Stops the loop if the number is **8** using `break`

#### **Task 3: Use a `switch` Statement**
Write a program that asks for a **day of the week** and prints:
- `"Weekend"` if it's **Saturday** or **Sunday**
- `"Weekday"` otherwise

---

### **Conclusion**
- **Conditionals (`if`, `switch`)** allow decision-making.
- **Loops (`for`, `range`)** control repetition.
- **Loop control (`break`, `continue`)** modifies loop execution.

---

Would you like me to provide the solution right away, or do you want to attempt it first? 🚀