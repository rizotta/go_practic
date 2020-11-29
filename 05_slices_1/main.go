package main

import "fmt"

func main() {
	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Create new slice + delete work days
	workDays := make([]string, len(weekDays)-2)
	weekDays = weekDays[copy(workDays, weekDays[:5]):]

	// Display
	fmt.Print("Weekends: ")
	fmt.Println(weekDays)
	fmt.Print("Working days: ")
	fmt.Println(workDays)
}
