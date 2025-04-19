package main

import (
	"encoding/json"
	"fmt"
)

// Struct to be encoded/decoded to/from JSON
type Person struct {
	Name    string `json:"name"`      // Tag to map "Name" field to JSON key "name"
	Age     int    `json:"age"`       // Tag to map "Age" field to JSON key "age"
	IsAdult bool   `json:"is_adult"`  // Tag to map "IsAdult" to "is_adult"
}

func main() {
	fmt.Println("LEARNING JSON IN GO:")

	// Create a Person struct instance
	person := Person{Name: "Deepak", Age: 24, IsAdult: true}
	fmt.Println("Original struct:", person)

	// ===== ENCODING: Struct to JSON (Marshalling) =====
	jsonData, err := json.Marshal(person) // convert struct to JSON bytes
	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}
	fmt.Println("JSON Data (as string):", string(jsonData)) // convert bytes to string for display

	// ===== DECODING: JSON to Struct (Unmarshalling) =====
	var personData Person
	err = json.Unmarshal(jsonData, &personData) // convert JSON bytes back to struct
	if err != nil {
		fmt.Println("Error while unmarshalling:", err)
		return
	}
	fmt.Println("Decoded Struct:", personData)
}

/*
======= Additional Notes =======
- `json.Marshal()` takes a struct and returns a byte slice (JSON format).
- `json.Unmarshal()` takes a byte slice (JSON) and fills a struct from it.
- Use **struct tags** like `json:"name"` to define how struct fields appear in JSON.
- Tags are **case-sensitive** in JSON, so using lowercase in tags ensures consistency with typical JSON style.

======= Output Example =======
Original struct: {Deepak 24 true}
JSON Data (as string): {"name":"Deepak","age":24,"is_adult":true}
Decoded Struct: {Deepak 24 true}

======= Common Use Cases =======
- API request/response parsing
- Config file loading
- Struct-to-JSON conversion for logging, storage, etc.
*/
