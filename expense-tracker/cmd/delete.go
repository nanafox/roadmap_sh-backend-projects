package cmd

import (
	"flag"
	"fmt"

	"github.com/nanafox/expense-tracker/internal/expense"
)

// deleteHandler handles the `delete command`
func deleteHandler(args []string) (err error) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	expenseId := deleteCmd.Int("id", 0, "The ID of the expense to delete")

	deleteCmd.Usage = func() {
		usageText := `USAGE: expense-tracker delete --id EXPENSE_ID

The 'delete' command removes an expense from the database.
It requires the ID of the expense to be deleted.

OPTIONS:
========
--id EXPENSE_ID          The unique ID of the expense to be deleted. This is the ID
                         that was assigned when the expense was created.

EXAMPLES:
=========
  expense-tracker delete --id 123  # This deletes the expense with ID 123 from the database.
`

		fmt.Println(usageText)
	}

	deleteCmd.Parse(args)

	if *expenseId == 0 {
		deleteCmd.Usage()
		return
	}

	err = expense.DeleteById(*expenseId)
	if err != nil {
		return err
	}

	fmt.Printf("Expense with ID %d deleted successfully\n", *expenseId)

	return nil
}
