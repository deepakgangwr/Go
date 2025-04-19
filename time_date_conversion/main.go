package main

import (
	"fmt"
	"time"
)

func main() {
	// "dd-mm-yyyy"
	// "yyyy/mm/dd"

	// Go's reference date for formatting is: 2006-01-02 15:04:05 (Mon Jan 2 15:04:05 MST 2006)
	// It uses this specific value to format custom layouts.

	currentTime := time.Now() // Gets the current local time
	fmt.Println("currentTime : ", currentTime)
	fmt.Printf("type of currentTime is: %T\n  ", currentTime)

	// Incorrect format like "dd-mm-yyyy" will give wrong result
	// Only use the Go layout reference: 02-01-2006, 15:04:05 etc.
	formatted := currentTime.Format("02-01-2006, Monday, 3:04:05 PM") // Formats using Go's layout
	fmt.Println("formatted time:", formatted)

	layout_str := "2006-01-02" // layout string for parsing date
	dateStr := "2025-04-16"    // date string in that format
	formatted_time, _ := time.Parse(layout_str, dateStr) // Parse returns time.Time and error
	fmt.Println("formatted time:", formatted_time)

	// Adding 2 days to current time
	new_date := currentTime.Add(2 * 24 * time.Hour) // Add duration (2 days)
	fmt.Println("new_date time: ", new_date)

	formatted_new_date := new_date.Format("2006/01/02 Monday ") // Format the new date
	fmt.Println("formatted new_date time: ", formatted_new_date)
}


// ===== Additional Notes on time package =====
// - `time.Now()` → returns current time
// - `time.Format(layout string)` → format using Go’s reference time layout
//     * 2006 - Year
//     * 01   - Month
//     * 02   - Day
//     * 15   - Hour (24hr)
//     * 03   - Hour (12hr)
//     * 04   - Minute
//     * 05   - Second
//     * PM   - AM/PM
//     * Monday - Day name
//
// - `time.Parse(layout, value)` → parses a string into a time.Time object
// - `Add(duration)` → adds time to a date. Ex: `2 * time.Hour`, `24 * time.Hour`
// - Durations: `time.Second`, `time.Minute`, `time.Hour`, `24*time.Hour` for a day

// time package reference: https://pkg.go.dev/time
