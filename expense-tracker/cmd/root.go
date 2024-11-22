package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Execute executes the command provided on the CLI. It does this by
// transferring the operation to the specialized function for the action
// requested by the user.
func Execute() (err error) {
	flag.Usage = usage

	if len(os.Args) < 2 {
		flag.Usage()
		return errors.New("Expected a command")
	}

	switch os.Args[1] {
	case "add":
		return addHandler(os.Args[2:])
	case "list":
		return listHandler(os.Args[2:])
	case "delete":
		return deleteHandler(os.Args[2:])
	case "update":
		return notImplementedError("update")
	case "summary":
		return summaryHandler(os.Args[2:])
	default:
		flag.Usage()
		return notImplementedError(os.Args[1])
	}
}

// usage prints the usage info for the expense-tracker CLI tool.
func usage() {
	usageText := `expense-tracker is a CLI expense expense-tracker

Usage:
  expense-tracker command [arguments]

  Implemented commands:
    add        Adds a new expense to the list of tracked expenses.
    list       Prints all the expenses available in the system.
    delete     Deletes an expense. CAUTION: This is not reversible.
    update     Update an an existing expense, useful when correcting mistakes.
    summary    Prints the sum of all expenses.

  For command-specific help, use "expense-tracker [command] --help"
`

	fmt.Println(usageText)
}

// notImplementedError returns an error when a command is not implement.
func notImplementedError(cmd string) (err error) {
	return fmt.Errorf("%s: not implemented", cmd)
}
