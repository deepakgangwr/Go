package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Struct to map JSON response from API
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`    // ID will be auto-filled by the API during POST
	Title     string `json:"title"` // Struct tag must match exact JSON key
	Completed bool   `json:"completed"`
}

// ========== READ (GET) ==========
func performGetRequest() {
	res, err := http.Get("https://dummyjson.com/todos/1")
	if err != nil {
		fmt.Println("Error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in response:", res.Status)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Fetched Todo Struct:", todo)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)
}

// ========== CREATE (POST) ==========
func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "DEEPAK",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	jsonReader := strings.NewReader(string(jsonData))
	myURL := "https://dummyjson.com/todos/add"
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("POST Response:", string(data))
}

// ========== UPDATE (PUT) ==========
func performUpdateRequest() {
	todo := Todo{
		UserId:    99,
		Title:     "Updated Title",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	jsonReader := strings.NewReader(string(jsonData))
	const myURL = "https://dummyjson.com/todos/1" // Make sure ID in URL is valid

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("PUT Response:", string(data))
}

// ========== DELETE ==========
func performDeleteRequest() {
	const myURL = "https://dummyjson.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("DELETE Response:", string(data))
}

// ========== Entry Point ==========
func main() {
	fmt.Println("Learning CRUD operation in Go:")

	// Uncomment the ones you want to run:
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}

/*
========= CRUD Summary =========
- GET (Read):    Use `http.Get()` to fetch data.
- POST (Create): Use `http.Post()` to create new data.
- PUT (Update):  Use `http.NewRequest("PUT", ...)` with `http.Client`.
- DELETE:        Use `http.NewRequest("DELETE", ...)` with `http.Client`.

========= Tips =========
- Always match struct tags to JSON keys (case-sensitive).
- Use `json.NewDecoder(res.Body).Decode(&obj)` for decoding.
- Use `defer res.Body.Close()` to avoid memory leaks.
*/
