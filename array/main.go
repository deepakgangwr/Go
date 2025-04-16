package main

import "fmt"

func main() {
	fmt.Println("array in go")

	var name[5]string
	name[0] = "deepak"
	name[2] = "Gangwar"

	fmt.Println(name)
    fmt.Printf("%q\n", name)
	// If you want to print a quoted string using %q,
	// you should use **fmt.Printf** instead of fmt.Println.

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("lenght of array is: ", len(nums))

	fmt.Println("value of name at 2nd index is : ", name[2])

	var price[5]int
	fmt.Println(price)
}

// in go , when we declare the array or slice. the elements are initialiazed to their zero values. 
// the zero value is the default value assiogned to variables of a certain type when no explicit value is provided
// for int it is 0
// for string ""
// for boolean is is false

// ==== Additional Notes on Arrays in Go ====
// - Arrays have **fixed length**, declared using [size]type (e.g., [5]int).
// - If you access an index beyond its length, it results in a **runtime panic**.
// - The size of the array is part of its type, so [5]int and [10]int are different types.
// - Arrays can be **initialized partially**; unassigned elements get their zero values.
// - You can also let Go infer the size using `[...]`, e.g., nums := [...]int{1, 2, 3}
// - Arrays can be compared using `==` if they are of the same type (same size and element type).
