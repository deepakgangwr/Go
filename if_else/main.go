package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(("x is greater than 5"))
	} else if x > 2 && x < 5 {
		fmt.Println("x is smaller than 5 but greater than 2")
	} else {
		fmt.Println("x is smaller than 3")
	}
}

// ====== Notes on Conditional Statements in Go ======
// - Curly braces `{}` are mandatory and must be on the same line as the condition.
// - Parentheses around conditions are optional. You can write:
//     if x > 5 { ... }  ✅
//     if (x > 5) { ... }  ✅ (still valid, but not idiomatic in Go)
// - You can declare and use a variable just for the condition:
//     if y := someFunc(); y > 10 { ... }
// - Logical operators: && (AND), || (OR), ! (NOT)
// - Comparison operators: >, <, >=, <=, ==, !=
