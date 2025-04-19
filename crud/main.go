package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to map JSON response from API
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"` // FIXED: Removed extra space in json tag
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD operation in Go:")

	// ===== READ Operation (GET Request) =====
	res, err := http.Get("https://dummyjson.com/todos/1") // fetch data of a single todo
	if err != nil {
		fmt.Println("Error getting:", err)
		return
	}
	defer res.Body.Close() // ensure the response body is closed after function ends

	// Check for non-200 status code
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response:", res.Status)
		return
	}

	// Option 1: Read raw data using io.ReadAll (commented for reference)
	/*
		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Println("Data:", string(data))
	*/

	// Option 2: Decode directly into struct (more efficient)
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// Output the todo
	fmt.Println("Fetched Todo Struct:", todo)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)

	// ===== Placeholder for other CRUD Operations =====
	// CREATE: Use http.Post() or http.NewRequest("POST", ...) with http.Client
	// UPDATE: Use http.NewRequest("PUT", ...) with JSON body and client.Do()
	// DELETE: Use http.NewRequest("DELETE", ...) and client.Do()

	// These are not implemented yet but follow this structure:
	/*
		// CREATE Example:
		payload := Todo{UserId: 1, Title: "New Task", Completed: false}
		body, _ := json.Marshal(payload)
		resp, err := http.Post("https://dummyjson.com/todos/add", "application/json", bytes.NewBuffer(body))

		// UPDATE Example:
		req, _ := http.NewRequest("PUT", "https://dummyjson.com/todos/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)

		// DELETE Example:
		req, _ := http.NewRequest("DELETE", "https://dummyjson.com/todos/1", nil)
		client := &http.Client{}
		resp, err := client.Do(req)
	*/
}

/*
========= Additional Notes =========
- `json:"title"`: JSON struct tags must match exact key names (case-sensitive, no spaces!).
- `http.Get()` performs a GET request and returns a response.
- `json.NewDecoder(res.Body).Decode(&struct)` decodes JSON directly from the response body.
- Always close response body with `defer res.Body.Close()`.
- For full CRUD, use `http.NewRequest()` with `http.Client` for methods other than GET and POST.
- Use `bytes.NewBuffer()` for request body in POST/PUT.

========= Sample Output =========
Fetched Todo Struct: {1 1 Learn Go true}
Title: Learn Go
Completed: true
*/

