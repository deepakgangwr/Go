package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",") // splits the string at commas
	fmt.Println(parts)                // [apple orange banana]

	data = "apple,orange,banana.DEEPaK"
	partss := strings.Split(data, ".") // splits the string at the period (.)
	fmt.Println(partss)                // [apple,orange,banana DEEPaK]

	str := "one two three four five four two two one"
	count := strings.Count(str, "two") // counts how many times "two" appears
	fmt.Println("count of two: ", count)

	str = "     Hello, Go!     "
	trimmed := strings.TrimSpace(str) // removes leading and trailing spaces
	fmt.Println("trimmed:", trimmed)

	str1 := "deepak"
	str2 := "gangwar"
	result := strings.Join([]string{"hello", str1, str2}, " ") // joins with space separator
	fmt.Println(result)
}

// ===== Additional Notes on strings package =====
// - strings.Split(str, sep)      → Splits a string into a slice based on separator
// - strings.Count(str, substr)   → Counts how many times substr appears in the string
// - strings.TrimSpace(str)       → Removes leading and trailing whitespace
// - strings.Join([]string, sep)  → Joins elements of a slice with a separator into one string

// Other Useful Functions:
// - strings.ToLower(str)         → Converts all letters to lowercase
// - strings.ToUpper(str)         → Converts all letters to uppercase
// - strings.Contains(str, sub)   → Returns true if substring exists
// - strings.HasPrefix(str, pre)  → Checks if string starts with prefix
// - strings.HasSuffix(str, suf)  → Checks if string ends with suffix
// - strings.ReplaceAll(str, old, new) → Replaces all occurrences of a substring

// strings package reference: https://pkg.go.dev/strings
