## **Lesson 4: Functions in Go**

### **Objectives:**
By the end of this lesson, you will:
- Understand how to declare and call functions in Go.
- Learn about function parameters and return values.
- Explore multiple return values.
- Compare Go functions with Python functions.
- Learn about named return values.

---

### **1. Functions in Go**
Functions in Go are similar to Python functions but require explicit type declarations for parameters and return values.

#### **Basic Function Declaration**
A simple function in Go:
```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello from Go!")
}

func main() {
    greet()
}
```
- `func greet()` declares a function named `greet`.
- The function is called in `main()`.

#### **Function with Parameters**
Unlike Python, Go requires explicit type definitions for each parameter.

Python:
```python
def add(x, y):
    return x + y
```
Go:
```go
func add(x int, y int) int {
    return x + y
}

func main() {
    sum := add(5, 7)
    fmt.Println("Sum:", sum)
}
```
- `x int, y int` declares two integer parameters.
- The return type `int` is explicitly declared after the parameter list.

---

### **2. Multiple Return Values**
Go allows functions to return multiple values.

#### **Example: Returning Two Values**
```go
func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)
}
```
- The return type `(int, int)` specifies two integers.
- Values are returned in order and can be assigned to multiple variables.

---

### **3. Named Return Values**
Go supports named return values, which act like declared variables.

```go
func rectangleDimensions(length int, width int) (area int, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return  // Explicit return with named values
}

func main() {
    a, p := rectangleDimensions(5, 3)
    fmt.Println("Area:", a, "Perimeter:", p)
}
```
- `area` and `perimeter` are named return variables.
- Simply using `return` returns their values.

---

### **4. Functions vs Python**
| Feature | Python | Go |
|---------|--------|----|
| Syntax | `def func_name():` | `func func_name() {}` |
| Type Hinting | Optional (`def add(x: int, y: int) -> int:`) | Required (`func add(x int, y int) int`) |
| Multiple Returns | `return a, b` | `return a, b` (Same approach) |
| Named Returns | No equivalent | Supported (`return` without arguments) |

### **Conclusion**
- Functions in Go require explicit types for parameters and return values.
- Multiple return values make Go functions more powerful.
- Named return values can simplify function definitions.