package main

import "fmt"

func main() {
	age := 20
	name := "deepak"
	height := 5.8
	fmt.Println("age: ", age, " height: ", height, " name: ", name)
	fmt.Println("hello world")

	fmt.Printf("age: %d\n", age)
	fmt.Printf("height: %.2f\n", height)
	fmt.Printf("name: %T\n", name)
	fmt.Printf("name: %s\n",name)
}
// In Go (fmt package), Println and Printf are used for outputting text, but they work slightly differently:

// Feature	fmt.Println()	fmt.Printf()
// Usage	Prints values with a space between them, and adds a newline (\n) automatically	Prints values according to a format string (you control how each value looks)
// Formatting	No custom formatting (just raw values)	Custom formatting with placeholders like %d, %f, %s, etc.
// Automatic Space	Yes	No
// Automatic Newline	Yes	No (unless you manually add \n)
// Example	fmt.Println("Hello", 123) → Hello 123\n	fmt.Printf("Hello %d\n", 123) → Hello 123\n
// ✅ Your Code:

// go
// Copy code
// fmt.Println("age: ", age, " height: ", height, " name: ", name)
// Println prints everything with spaces automatically.

// Example output:

// less
// Copy code
// age: 20 height: 5.8 name: deepak
// go
// Copy code
// fmt.Printf("age: %d\n", age)
// fmt.Printf("height: %.2f\n", height)
// fmt.Printf("name: %T\n", name)
// Printf gives formatted output:

// %d → integer (decimal)

// %.2f → float with 2 decimal places

// %T → prints the type of the variable (string, int, etc.)

// Example Output:

// vbnet
// Copy code
// age: 20
// height: 5.80
// name: string
// ✅ Summary:

// If you want	Use
// Just to quickly print values	Println()
// To control how values look (formatting numbers, strings, types)	Printf()
// ⚡ Extra Tip:

// There's also fmt.Print() → prints without any space or newline automatically.

// Println = print + newline

// Printf = print with format (you control newline)