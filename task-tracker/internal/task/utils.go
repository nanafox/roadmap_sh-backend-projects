package task

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// nextId returns the next ID to be used for the task being created.
func nextId() (Id int) {
	return len(tempStorage) + 1
}

// itoa converts a number to a string.
func itoa(n int) string {
	return strconv.Itoa(n)
}

// createFile creates the tasks file if it does not already exist.
func createFile(filename string) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return
	}
}

// tasksPrinter prints the tasks provided.
func tasksPrinter(tasks []Task, status string) {
	header := setHeader(status)

	fmt.Println(header)

	for _, task := range tasks {
		printDivider()

		fmt.Println("ID:", task.Id)
		fmt.Println("Status:", task.Status)
		fmt.Println("Description:", task.Description)

		printDivider()
	}
}

// printDivider prints a simple divider between outputs
func printDivider() {
	fmt.Printf("\n%s\n", strings.Repeat("=", 54))
}

// setHeader sets the header for printing tasks.
//
// The header is set based on the status of the status that is being retrieved.
// This helps the user to know what tasks they are looking at.
func setHeader(status string) (header string) {
	switch status {
	case "":
		header = "ALL TASKS"
	case "todo":
		header = "TODO TASKS"
	case "in-progress":
		header = "TASKS IN PROGRESS"
	case "done":
		header = "COMPLETED TASKS"
	}

	return
}
