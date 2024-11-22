package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
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

// GetAll returns all the expenses in the storage.
//
// It sorts the expenses by CreatedAt and limits the results to the specified
// number of expenses (limit).
// Parameters:
//
//	limit (int): The number of expenses to retrieve
func GetAll(limit int) (expenses []Expense) {
	// Convert map to slice
	expenses = make([]Expense, 0, len(expenseStorage.Expenses))
	for _, expense := range expenseStorage.Expenses {
		expenses = append(expenses, expense)
	}

	// Sort by CreatedAt
	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].CreatedAt.Before(expenses[j].CreatedAt)
	})

	// Apply limit if needed
	if limit > 0 && limit < len(expenses) {
		expenses = expenses[:limit]
	}

	return expenses
}

// getById retrieves an expense by its ID from the ExpenseStorage.
func getById(id int) (expense Expense, err error) {
	expense, found := expenseStorage.Expenses[id]

	if !found {
		err = fmt.Errorf("Expense with ID %d not found", id)
	}

	return
}

// DeleteById removes the expense by its ID.
func DeleteById(id int) (err error) {
	_, exists := expenseStorage.Expenses[id]
	if !exists {
		return fmt.Errorf("Expense with ID %d not found", id)
	}

	delete(expenseStorage.Expenses, id)

	return persistExpenses()
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
