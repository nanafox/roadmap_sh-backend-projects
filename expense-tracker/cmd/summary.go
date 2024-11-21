package cmd

import (
	"flag"
	"fmt"
	"time"

	"github.com/nanafox/expense-tracker/internal/expense"
)

// summaryHandler handles the `summary` command.
//
// The `summary` command prints the total amount of money spent on expenses. It
// does this in two ways.
//
//  1. It prints a summary of all expenses.
//  2. It prints the summary of expenses for a specific month of the current
//     year. This requires the `--month` option.
func summaryHandler(args []string) (err error) {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	summaryCmd.Usage = func() {
		usageText := `USAGE: expense-tracker summary [--month MONTH_NUMBER]

The 'summary' command prints the total amount of money spent on expenses. 
It does this in one of two ways.

  1. It prints a summary of all expenses.
  2. It prints the summary of expenses for a specific month of the current
     year. This requires the '--month' option.

OPTIONS:
========
--month MONTH_NUMBER    The month of the current year to retrieve expenses for.
                        This number is the numeric value of the month.

EXAMPLES:
=========
  expense-tracker summary  # This prints all the expenses available in the database.
    -OR-
  expense-tracker summary --month 8  # This prints the expenses for August in the current year.
`

		fmt.Println(usageText)
	}
	month := summaryCmd.Int(
		"month",
		0,
		"The month to retrieve summary of expenses from.",
	)

	expenses := expense.GetAll(expense.NumberOfExpenses())

	summaryCmd.Parse(args)
	total := 0.0

	if *month != 0 {
		for _, expense := range expenses {
			if int(expense.CreatedAt.Month()) == *month {
				total += expense.Amount
			}
		}

		fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(*month), total)
	} else {
		for _, expense := range expenses {
			total += expense.Amount
		}

		fmt.Printf("Total expenses: $%.2f\n", total)
	}

	return
}
