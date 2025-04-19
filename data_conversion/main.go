package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 5
	fmt.Println(num)
	fmt.Printf("type of num is %T\n", num)

	// num = num + 1.23 // throws error (type mismatch)
	var data float64 = float64(num) // int to float
	fmt.Println(data)
	fmt.Printf("type of data is %T\n", data)

	data = data + 1.21
	fmt.Println(data)

	// ===== int to string =====
	num = 123
	str := strconv.Itoa(num) // Itoa = Int to ASCII
	fmt.Println(str)
	fmt.Printf("type of str is %T\n", str)

	// ===== string to int =====
	num_string := "1234"
	// num_string = num_string + 123 // throws error
	num_int, _ := strconv.Atoi(num_string) // Atoi = ASCII to Int
	fmt.Println(num_int)
	num_int = num_int + 123
	fmt.Println(num_int)
	fmt.Printf("type of num_int is %T\n", num_int)

	// ===== string to float =====
	num_string = "123.4"
	num_float, _ := strconv.ParseFloat(num_string, 64) // Atoi won't work here, use ParseFloat
	fmt.Println(num_float)
	num_float = num_float + 1.6
	fmt.Println(num_float)
	fmt.Printf("type of num_float is %T\n", num_float)
}

// ===== Additional Notes =====
// - strconv.Itoa(int) → string
// - strconv.Atoi(string) → int
// - strconv.ParseFloat(string, 64) → float64
// - All conversions return 2 values: (result, error), use `_` if you want to ignore error
// - Direct operations between int and float are not allowed in Go (need conversion)
// - You cannot convert between string and number directly using type casting; use strconv
// - strconv is part of Go standard library (docs: https://pkg.go.dev/strconv)
