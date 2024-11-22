package cmd

import (
	"flag"
	"fmt"

	"github.com/nanafox/expense-tracker/internal/expense"
)

// updateUsage prints the the usage info for the `update` command.
func updateUsage() {
	usageText := `USAGE: expense-tracker update --id EXPENSE_ID --description DESCRIPTION --amount AMOUNT

The 'update' command modifies the details of an existing expense in the database.
It requires the ID of the expense to update and allows updating the description, amount, or both.

OPTIONS:
========
--id EXPENSE_ID             The unique ID of the expense to be updated. This ID identifies the specific expense.
--description DESCRIPTION   The new description for the expense. If omitted, the description remains unchanged.
--amount AMOUNT             The new amount for the expense. If omitted, the amount remains unchanged.

EXAMPLES:
=========
  expense-tracker update --id 123 --description "Lunch with team" --amount 45.50
    # Updates the expense with ID 123, setting the description to "Lunch with team" and the amount to 45.50.

  expense-tracker update --id 456 --description "Office supplies"
    # Updates only the description of the expense with ID 456.

  expense-tracker update --id 789 --amount 25.00
    # Updates only the amount of the expense with ID 789.
`

	fmt.Println(usageText)
}

// updateHandler handles the `update` command.
func updateHandler(args []string) (err error) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	expenseId := updateCmd.Int("id", 0, "The ID of the expense to update")
	description := updateCmd.String("description", "", "The new description for the expense")
	amount := updateCmd.Float64("amount", 0.0, "The new amount to set.")

	updateCmd.Usage = updateUsage

	updateCmd.Parse(args)

	if *expenseId == 0 {
		updateCmd.Usage()
		return
	}

	err = expense.UpdateById(*expenseId, *description, *amount)
	if err != nil {
		return err
	}

	fmt.Println("Expense updated successfully")

	return nil
}
