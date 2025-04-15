package main

import (
	"fmt"
// 	"mylearning/myutil"
// 	"mylearning/variables"
)

func main() {
	fmt.Println("Hello, Go!")

	// Using variables from variables.go
	// fmt.Println(variables.Name)
	// fmt.Println(variables.Version)
	// fmt.Println(variables.Money)
	// fmt.Println(variables.Currency)
	// fmt.Println(variables.Dimension)
	// fmt.Println(variables.Flo)
	// fmt.Println(variables.Flag)
	// fmt.Println(variables.PiConst)
	// fmt.Println(variables.Person)
	// fmt.Println(variables.Public)

	// Note: private (small p) variable from variables.go is not accessible here
	// because in Go, anything starting with a lowercase letter is unexported (only visible within the same package)

	// Using functions from myutil.go
	// myutil.PrintMessage("Hello from myutil package!")
	// myutil.Test()

	// myutil.messagePrinter() // ❌ It shows "messagePrinter" as undeclared
	// because it starts with a small letter, so it is private (unexported)
}
