package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

func main() {
	// Read task data from README.md
	data, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println("Error reading README.md:", err)
		return
	}

	// Regex for tasks (adjust if needed)
	totalTasksRegex := regexp.MustCompile(`\[.\]`)
	completedTasksRegex := regexp.MustCompile(`\[x\]`)

	totalTasks := len(totalTasksRegex.FindAllString(string(data), -1))
	completedTasks := len(completedTasksRegex.FindAllString(string(data), -1))
	tasksLeft := totalTasks - completedTasks

	percentComplete := float64(completedTasks) / float64(totalTasks) * 100

	now := time.Now()
	targetDate := time.Date(now.Year(), time.June, 30, 0, 0, 0, 0, time.UTC)
	daysUntil := int(targetDate.Sub(now).Hours() / 24)

	dayPercentComplete := float64(tasksLeft) / (float64(daysUntil) / 7)

	// Output results
	if tasksLeft == 0 {
		fmt.Println("All Done Successfully!")
	} else {
		fmt.Println("Completed:", completedTasks)
		fmt.Printf("Remaining: %d days\n", tasksLeft)
		fmt.Printf("Percent Complete: %.2f%%\n", percentComplete)
		fmt.Println("Days Remaining:", daysUntil)
		fmt.Printf("Minimum Required Average Task Completion Rate per week: %.3f\n", dayPercentComplete)
	}
}
