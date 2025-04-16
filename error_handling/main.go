package main

import "fmt"

// Error handling using multiple return values
// func divide(a,b float64) (float64,error) {
// 	if(b==0){
// 		return 0,fmt.Errorf("denominator must not be zero")
// 	}
// 	return a/b,nil
// }

// Custom version using string instead of built-in error type
func divide(a,b float64) (float64,string) {
	if(b==0){
		return 0,"denominator must not be zero"
	}
	return a/b,"nil"
}

func main() {
	fmt.Println("error handling")

	// Using Go's built-in error interface
	// ans, err:=divide(10,0)
	// if err !=nil{
	// 	fmt.Println("error handling")
	// }
	// fmt.Println(ans)

	ans, _ :=divide(10,2) // _ is used to ignore the element returned by function that we don't need
	fmt.Println(ans)
}

// === Additional Notes on Error Handling in Go ===
// - The standard practice in Go is to return an `error` type as the second return value.
// - `fmt.Errorf("message")` creates a new error with a formatted message.
// - Always check if the error is `nil` before using the returned result.
// - `if err != nil { ... }` is the common idiom for error checking.
// - Using `string` for errors (like above) is possible but not idiomatic in production Go code.
// - You can use `errors.New("message")` or `fmt.Errorf(...)` from the `errors` and `fmt` packages for better error management.
