package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"time"
)

var tempStorage map[string]Expense

var nextAvailableId int

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

	tempStorage = make(map[string]Expense)
	if len(file) != 0 {
		var fileContent ExpenseStorage

		err = json.Unmarshal(file, &fileContent)
		if err != nil {
			log.Fatal(err)
		}

		tempStorage = fileContent.Expenses
		nextAvailableId = fileContent.Metadata.NextId
	} else {
		nextAvailableId = 1
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

	tempStorage[itoa(expenseId)] = *expense

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
	expenseStorage := ExpenseStorage{
		Metadata: struct {
			NextId int `json:"nextId"`
		}{NextId: nextAvailableId}, // Using the global nextAvailableId value
		Expenses: tempStorage,
	}

	// Serialize the data
	jsonData, err := json.MarshalIndent(
		expenseStorage, "", "  ",
	)
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

// getAll returns all the expenses in the storage.
//
// When the status is a non-empty string, it is used to filter the expenses returned
// based on its status. The list of accepted expense statuses are as follows:
//
// todo => For expenses not yet started.
// in-progress => For expenses that has been started but not completed.
// done => For completed expenses.
func getAll() (expenses []Expense) {
	return slices.Collect(maps.Values(tempStorage))
}

// getById retrieves the expense by its ID.
func getById(id int) (expense Expense, err error) {
	strId := itoa(id)
	expense, found := tempStorage[strId]

	if !found {
		err = fmt.Errorf("Expense with ID %d not found", id)
	}

	return
}

// nextId returns the next ID to be used for the expense being created.
func nextId() (Id int) {
	Id = nextAvailableId
	nextAvailableId++
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
