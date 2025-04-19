package main

import "fmt"

func main() {
	fmt.Println("Starting of the program")

	data := 5

	defer fmt.Println(data + 5)               // 2nd defer - this will execute **first**
	defer fmt.Println("Middle of the program") // 1st defer - this will execute **second**

	fmt.Println("End of the program")
}

// ===== Notes on defer =====
// - `defer` is used to delay the execution of a function until the surrounding function returns.
// - Commonly used for cleanup tasks (closing files, unlocking resources, etc.).
// - Deferred calls are executed in **LIFO (Last In, First Out)** order.

// In this example:
// Output will be:
// Starting of the program
// End of the program
// Middle of the program     <-- second defer (called first)
// 10                        <-- first defer (called last)

// The value of variables used in defer statements are **evaluated at the time the defer is declared**.
// So `data + 5` evaluates to 10 immediately, not when the function ends.
