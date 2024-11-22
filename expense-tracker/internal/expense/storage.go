package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var expenseStorage ExpenseStorage

var filename string = os.Getenv("HOME") + "/.expenses.json"

// init initializes the storage for storing the expenses.
func init() {
	err := createFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	expenseStorage.Expenses = make(map[int]Expense)
	if len(file) != 0 {
		err = json.Unmarshal(file, &expenseStorage)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		expenseStorage.Metadata.NextId = 1
	}
}

// Save persists the expenses to the storage.
func (expense *Expense) Save(options map[string]bool) (expenseId int, err error) {
	currentTime := time.Now()

	if options != nil && options["update"] {
		expense.UpdatedAt = currentTime
	} else {
		expense.Id = nextId()
		expense.CreatedAt = currentTime
		expense.UpdatedAt = currentTime
	}

	expenseId = expense.Id

	expenseStorage.Expenses[expenseId] = *expense

	// Persist the expense data and metadata
	err = persistExpenses()
	if err != nil {
		return 0, err
	}

	return
}

// persistExpenses serializes the current expense data and metadata, and writes it
// to the storage file.
func persistExpenses() error {
	// Create the expense storage structure
	jsonData, err := json.MarshalIndent(expenseStorage, "", "  ")
	if err != nil {
		return fmt.Errorf("persist: error serializing expenses: %w", err)
	}

	// Write to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("persist: error writing to storage file: %w", err)
	}

	return nil
}

// nextId returns the next ID to be used for the expense being created.
func nextId() (id int) {
	id = expenseStorage.Metadata.NextId
	expenseStorage.Metadata.NextId++
	return
}

// itoa converts a number to a string.
func itoa(n int) string {
	return strconv.Itoa(n)
}

// createFile creates the expenses file if it does not already exist.
func createFile(filename string) (err error) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}

	return
}

// NumberOfExpenses returns the number of expenses in the database.
func NumberOfExpenses() (numOfExpenses int) {
	return len(expenseStorage.Expenses)
}
