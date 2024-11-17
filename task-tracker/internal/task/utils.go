package task

import (
	"errors"
	"log"
	"os"
	"strconv"
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
