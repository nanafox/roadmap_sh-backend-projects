package main

import (
	"fmt"
	"os"

	"github.com/nanafox/expense-tracker/cmd"
)

// main is the entry point to the expense-tracker CLI program.
func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
