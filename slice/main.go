package main

import "fmt"

func main() {
	fmt.Println("study slice")

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Printf("%T\n", nums) // %T to print the data type.
	fmt.Println("lenght of slice : ", len(nums))

	nums = append(nums, 6, 7, 8, 9, 10)
	fmt.Println(nums)
	fmt.Println("lenght of slice : ", len(nums))

	fmt.Println("slice:", nums)
	fmt.Println("lenght", len(nums))
	fmt.Println("capacity", cap(nums))

	num := make([]int, 3, 5) // variable len and cap
	fmt.Println("slice:", num)
	fmt.Println("lenght", len(num))
	fmt.Println("capacity", cap(num))

	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6)

	fmt.Println("slice:", num)
	fmt.Println("lenght", len(num))
	fmt.Println("capacity", cap(num)) // cap *= 2
}

// if we want to initialize the empty slice than use make function or add {} like nums:=[]int{}
// nums := make([]int, 0)
// nums := []int{}

// ====== Additional Notes on Slices in Go ======
// - Slices are dynamically-sized, unlike arrays.
// - `len(slice)` gives the current number of elements.
// - `cap(slice)` shows the underlying array's capacity.
// - When appending beyond capacity, Go automatically allocates a new underlying array with double the capacity.
// - `make([]int, length, capacity)` lets you set initial length and capacity manually.
// - Slices are reference types — assigning one to another points to the same underlying array.
// - Use `append()` to grow slices, even if they're empty.
// - `append()` returns a new slice — always assign it back if needed (e.g., s = append(s, 10))

