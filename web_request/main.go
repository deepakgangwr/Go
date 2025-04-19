package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web services (GET request)")

	// Sending a GET request to a public API
	res, err := http.Get("https://dummyjson.com/todos")
	if err != nil {
		fmt.Println("error getting response: ", err)
		return
	}
	defer res.Body.Close() // important to close response body

	fmt.Printf("type of res: %T\n", res) // *http.Response

	// Reading response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	// `data` is a byte slice, so print it as a string to see readable output
	// fmt.Println("response (bytes):", data)
	fmt.Println("response (string):", string(data)) // pretty JSON string
}

// ===== Additional Notes =====
// - This makes a simple HTTP GET request to a public API.
// - Always check for errors while making requests or reading response bodies.
// - `defer res.Body.Close()` ensures the response body is properly closed (prevents leaks).
// - `io.ReadAll()` reads the entire body; good for small to medium-sized responses.
// - Use `string(data)` to convert byte slice to readable string (usually JSON).
// - To parse JSON into Go structs, use `encoding/json` package with `json.Unmarshal()` (optional next step).
