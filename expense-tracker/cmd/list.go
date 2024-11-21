package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/nanafox/expense-tracker/internal/expense"
)

// listUsage prints the usage information for the list command.
func listUsage() {
	usageText := `USAGE: expense-tracker list [--limit=NUMBER]

The 'list' command prints all the expenses available in the storage. By default,
it returns all the expenses but it provides the '--limit' option to allow you
to specify how many expenses you want to return at a time.

OPTIONS:
========
--limit=NUMBER  The number of expenses to print. Default limit is total number
                of available expenses at the time the command is executed.

EXAMPLES:
=========
  expense-tracker list  # This will print all expenses
    -OR-
  expense-tracker list 2  # This will print the first expenses
`
	fmt.Println(usageText)
}

// listHandler handle the `list` command.
//
// This list command prints all the available expenses with an option to limit
// the number of expenses printed by specifying a number to the `--limit`
// option.
func listHandler(args []string) (err error) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.Usage = listUsage

	limit := listCmd.Int("limit", expense.NumberOfExpenses(), "The number of expenses to retrieve")

	listCmd.Parse(args)

	if *limit < 1 {
		return errors.New("list: limit must be greater than or equal to 1")
	}
	expenses := expense.GetAll(*limit)

	printExpenses(expenses)
	return
}

// getMaxDescriptionLength returns the maximum length of the description
func getMaxDescriptionLength(expenses []expense.Expense) (maxDescriptionLength int) {
	maxDescriptionLength = len("Description")

	for _, expense := range expenses {
		if len(expense.Description) > maxDescriptionLength {
			maxDescriptionLength = len(expense.Description)
		}
	}

	return
}

// printExpenses prints all the expenses provided in the slice of expenses.
func printExpenses(expenses []expense.Expense) {
	maxDescriptionLength := getMaxDescriptionLength(expenses)

	// Print the header
	fmt.Printf(
		"# %-3s %-10s %-*s %10s\n",
		"ID",
		"Date",
		maxDescriptionLength,
		"Description",
		"Amount",
	)

	// Print each expense
	for _, expense := range expenses {
		date := expense.CreatedAt
		formattedDate := fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
		fmt.Printf(
			"# %-3d %-10s %-*s %5s%.2f\n",
			expense.Id,
			formattedDate,
			maxDescriptionLength,
			expense.Description,
			"$",
			expense.Amount,
		)
	}
}
