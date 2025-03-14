## Variables and Constants Exercise
1. Declare a `var` integer named `counter` and set it to `100`.
2. Declare a constant named `appName` with the value `"GoApp"`.
3. Use `:=` to declare a boolean variable `isRunning` set to `true`.
4. Print all values using `fmt.Println`.

**Try it out:**
```go
package main

import "fmt"

func main() {
    var counter int = 100
    const appName = "GoApp"
    isRunning := true

    fmt.Println("Counter:", counter)
    fmt.Println("Application:", appName)
    fmt.Println("Is Running?", isRunning)
}
```