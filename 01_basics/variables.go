package main

import "fmt"

func main() {
    // Declare a variable
    var name string = "Go Learner"
    fmt.Println("Hello,", name)

    // Short declaration (type inferred)
    age := 78
    fmt.Println("Your age is:", age)

    // Declare a constant
    const pi float64 = 3.14159
    fmt.Println("Value of pi:", pi)

    // Multiple variables declarations
    var x, y, z int = 1, 2, 3
    fmt.Println("Multiple variables:", x, y, z)

    // Using a variable
    greeting := "Welcome to Go!"
    fmt.Println(greeting)
}