package cmd

import (
	"flag"
	"fmt"

	"github.com/nanafox/expense-tracker/internal/expense"
)

// addHandler handles the add command which is used to add new expenses to the
// expenses database.
//
// The function expects a slice of command strings which is essentially the
// `description` of the expense and the `amount` of the expense. Without these
// values, the function fails and returns an error. It also shows the usage
// information for the command.
//
// Parameters:
//
//	args([]string): A slice of command line arguments which represent the
//	                description and amount of the new expense.
//
// Returns:
//
//	(err: error): An error is returned when something goes wrong, nil otherwise.
func addHandler(args []string) (err error) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.Usage = func() {
		usageText := `USAGE: expense-tracker add --description TEXT --amount AMOUNT

The 'add' command gives you an interface to add new expense
to your list of tracked expenses.

OPTIONS:
========
  --description=TEXT   The description of the expense you want to add
  --amount=AMOUNT     The amount of money spent on an expense. Accepts decimals.

EXAMPLES:
=========
  expense-tracker add --description "Lunch" --amount 20.99
    -OR-
  expense-tracker add --description="Lunch" --amount=20.99
`

		fmt.Println(usageText)
	}

	description := addCmd.String("description", "", "The description of the expense")
	amount := addCmd.Float64("amount", 0, "The amount spent on this expense")

	addCmd.Parse(args)

	if description == nil || amount == nil {
		addCmd.Usage()
		fmt.Println(args)
		return
	}

	newExpense := expense.Expense{
		Description: *description,
		Amount:      *amount,
	}

	expenseId, err := newExpense.Save(nil)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", expenseId)
	return nil
}
