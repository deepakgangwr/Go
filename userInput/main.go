package main

import (
	"bufio"  // Import bufio package for buffered I/O
	"fmt"
	"os"     // Import os package to access standard input (keyboard)
)

func main() {
	fmt.Println("What's your name?")

	// Create a new buffered reader that reads from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Read input from the user until they press ENTER ('\n')
	fullName, _ := reader.ReadString('\n')  
	// ReadString returns two values: (string, error)
	// _ is used to ignore the error for now (good to handle it properly later)

	fmt.Println("Hello, Mr.", fullName)
}
