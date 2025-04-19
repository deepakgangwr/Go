package main

import "fmt"

func main() {
	// name -> grade
	studentGrade := make(map[string]int)

	studentGrade["A"] = 95
	studentGrade["B"] = 10
	studentGrade["C"] = 33
	studentGrade["D"] = 100
	studentGrade["E"] = 59

	fmt.Println("marks of D: ", studentGrade["D"])
	fmt.Println("marks of C: ", studentGrade["C"])

	// updating value
	studentGrade["C"] = 100
	fmt.Println("marks of C after update: ", studentGrade["C"])

	// deleting a key
	delete(studentGrade, "C")
	fmt.Println("marks of C after deletion: ", studentGrade["C"]) // returns zero value

	// checking if a key exists
	grades, exists := studentGrade["R"]
	fmt.Println("marks of R: ", grades)       // prints 0 (default int)
	fmt.Println("R exists: ", exists)         // prints false

	// iterating over map
	for index, value := range studentGrade {
		fmt.Printf("key is %s and marks is %d\n", index, value) // unordered output
	}


	person := map[string]int{
		"a":90,
		"b":50,
		"c":60,
	}
	for index, value := range person {
		fmt.Printf("key is %s and marks is %d\n", index, value) // unordered output
	}
}

// ===== Additional Notes =====
// - Maps in Go are unordered, like unordered_map in C++.
// - If you access a non-existent key, it returns the zero value (0 for int, "" for string, etc.).
// - Use the `value, exists := map[key]` pattern to check if a key is present.
// - You can also use literals to create a map:
//     student := map[string]int{"A": 90, "B": 80}
// - The `make(map[keyType]valueType)` function creates an empty map.
// - Deleting a non-existent key does not throw an error.
