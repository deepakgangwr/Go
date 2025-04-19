package main

import "fmt"

// defining struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

// nested struct
type Company struct {
	employeeID int
	role       string
	details    Person
}

func main() {
	// ===== Empty Struct and Assignment =====
	var Deepak Person
	fmt.Println(Deepak) // prints zero values

	Deepak.firstName = "Deepak"
	Deepak.lastName = "Gangwar"
	Deepak.age = 22
	fmt.Println("Updated Deepak:", Deepak)

	// ===== Struct Literal with Field Names =====
	person1 := Person{
		firstName: "Kaushal",
		lastName:  "Gangwar",
		age:       25,
	}
	fmt.Println("Person1:", person1)

	// ===== Using new keyword (returns pointer) =====
	var person2 = new(Person)
	person2.firstName = "Akansha"
	person2.lastName = "Raman"
	person2.age = 26
	fmt.Println("Pointer to struct:", person2)
	fmt.Println("Access field via pointer:", person2.firstName)

	// ===== Struct inside Struct Example =====
	employee := Company{
		employeeID: 101,
		role:       "Software Engineer",
		details: Person{
			firstName: "Shreya",
			lastName:  "Sharma",
			age:       28,
		},
	}
	fmt.Println("Employee Details:")
	fmt.Println("ID:", employee.employeeID)
	fmt.Println("Role:", employee.role)
	fmt.Println("Name:", employee.details.firstName, employee.details.lastName)
	fmt.Println("Age:", employee.details.age)
}

// ===== Additional Notes =====
// - Structs group multiple fields under one name.
// - Fields can be of different types.
// - If a field is not initialized, it gets the zero value:
//     int → 0, string → "", bool → false
// - You can access fields directly using dot notation.
// - Structs can be nested inside other structs.
// - `new(Type)` returns a pointer to a zero-valued instance.
// - Go also supports anonymous structs for quick use:
//     emp := struct{name string; age int}{"Raj", 30}
