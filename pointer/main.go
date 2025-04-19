package main

import "fmt"

// function to modify value by reference
func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	num := 5
	// var ptr *int
	// ptr = &num
	ptr := &num
	fmt.Println(num)   // prints: 5
	fmt.Println(ptr)   // prints: address of num
	fmt.Println(*ptr)  // prints: Deepak (value at address)

	// default pointer value is nil
	var pointer *int
	if pointer == nil {
		fmt.Println("pointer is not assigned") // this will print
	}

	// ===== Modify integer using pointer =====
	value := 10
	fmt.Println("value contains:", value)
	modifyValueByReference(&value)
	fmt.Println("value after modifyValueByReference:", value)
}

// ===== Additional Notes =====
// - A pointer holds the **memory address** of a variable.
// - Use `&` to get the address of a variable.
// - Use `*` to dereference the pointer and access/change the value.
// - Pointers in Go are type-safe: `*int` points to an `int`, `*string` to a `string`, etc.
// - Default value of an uninitialized pointer is `nil`.
// - Functions can use pointers to modify variables outside their local scope.
