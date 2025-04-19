package main

import (
	"fmt"
	"os"
	// Required for io.WriteString and io.EOF
)

func main() {
	// ===== Create and Write to a File =====
	/*
		file, err := os.Create("example.txt") // creates file or overwrites if exists
		if err != nil {
			fmt.Println("error while creating file:", err)
			return
		}
		defer file.Close() // important to release the file resource

		content := "hello file is created by deepak"
		byteWritten, writeErr := io.WriteString(file, content + "\n") // "\n" moves cursor to the next line

		fmt.Println("bytes written: ", byteWritten)
		if writeErr != nil {
			fmt.Println("error while writing to file:", writeErr)
			return
		}
		fmt.Println("file successfully created and written")
	*/

	// ===== Method 1: Read file using buffer (good for large files) =====
	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("error while opening the file:", err)
			return
		}
		defer file.Close()

		buffer := make([]byte, 1024) // buffer of 1KB

		for {
			n, err := file.Read(buffer)
			if err == io.EOF { // end of file
				break
			}
			if err != nil {
				fmt.Println("error reading the file:", err)
				return
			}

			fmt.Println(string(buffer[:n])) // only print read portion
		}
	*/

	// ===== Method 2: Read entire file using os.ReadFile (good for small files) =====

	content, err := os.ReadFile("example.txt") // was ioutil.ReadFile before Go 1.16
	if err != nil {
		fmt.Println("error while reading the file:", err)
		return
	}
	fmt.Println(string(content))
}

// ===== Additional Notes =====
// - `defer file.Close()` is critical to release system resources.
// - `io.WriteString()` is used to write a string into a file.
// - `os.Create()` creates or truncates a file; `os.Open()` opens in read-only mode.
// - `os.ReadFile()` loads the full file content at once — not ideal for huge files.
// - `os.Open()` + buffer loop is better for large file reading (chunk-wise).
// - Always check for `io.EOF` when reading files in chunks.
