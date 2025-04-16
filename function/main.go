package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

// Basic function with parameters and return type
func add(a int, b int) int {
	return a + b
}

// You can omit repeated type declarations
// func add(a, b, c, d, e int) — no need to add `int` for all, just add it at the last

// Named return values — useful for documentation and readability
// func add(a, b int) (result int) {
// 	result = a + b
// 	return
// }

func main() {
	fmt.Println("function")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println(ans)
}

// You can't start { } from next line after function declaration,
// you have to start from the same line — it gives an error.
// func ()  // {
// } — this will result in a syntax error

// === Additional Notes on Functions in Go ===
// - Functions are first-class citizens in Go: you can assign them to variables, pass them as arguments, and return them from other functions.
// - You can use variadic functions (like `func sum(nums ...int)`) to accept any number of arguments.
// - Functions can return multiple values, commonly used with `(result, err)` patterns.
// - Anonymous functions and closures are also supported using `func()` without a name.
