package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // skip the  loop when i equals 5
		}
		fmt.Println("i =", i)
	}

	// ===== Infinite loop with break =====
	counter :=1
	for {
		fmt.Println("This is an infinite loop")
		counter++
		if counter == 5{
			break // remove this to make it truly infinite
		}
	}

	// Corrected for-range loop
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("%d %d\n", index, value) // fixed parentheses in Printf
	}

	data :="hello, world!"
	for index, char := range data{
		fmt.Printf("%d %c\n",index,char)
	}

}
