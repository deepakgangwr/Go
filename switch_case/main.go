package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wedneday")
	default:
		fmt.Println("unknown day")
	}

	// Example: Grouping months using multiple values in case
	month := 6

	switch month {
	case 3, 4, 5:
		fmt.Println("Summer Season")
	case 6, 7, 8:
		fmt.Println("Rainy Season")
	case 11, 12, 1, 2:
		fmt.Println("Winter Season")
	default:
		fmt.Println("Invalid or Transition Month")
	}


// ===== switch with no expression (acts like switch true) =====
temp := 28

switch {
case temp >= 35:
	fmt.Println("Very Hot")
case temp >= 25 && temp < 35:
	fmt.Println("Pleasant Weather")
case temp >= 15 && temp < 25:
	fmt.Println("Cool")
default:
	fmt.Println("Cold")
}
}


// ====== Notes on switch in Go ======
// - No need to write `break` in Go; it auto breaks after each matching case.
// - Multiple values in a case can be separated using commas: case 1, 2, 3:
// - Use `fallthrough` to intentionally continue to the next case.
// - You can switch on types like int, string, bool, etc.
// - `switch` can be used without an expression as `switch true` for condition-based logic.
 