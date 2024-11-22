package expense

import (
	"fmt"
	"sort"
)

// UpdateById updates an expense by its ID.
func UpdateById(id int, description string, amount float64) (err error) {
	expense, err := getById(id)
	if err != nil {
		return err
	}

	if description != "" {
		expense.Description = description
	}

	if amount > 0.0 {
		expense.Amount = amount
	}
	_, err = expense.Save(map[string]bool{"update": true})
	if err != nil {
		return err
	}

	return nil
}

// DeleteById removes the expense by its ID.
func DeleteById(id int) (err error) {
	expense, err := getById(id)
	if err != nil {
		return err
	}

	delete(expenseStorage.Expenses, expense.Id)

	return persistExpenses()
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
