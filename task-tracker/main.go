package main

import (
	"fmt"
	"os"

	"github.com/nanafox/task-tracker/cmd"
)

// TODO: Include a good help system for task-cli so users know what to expect
// when using it.

// main provides the entry to the task-cli program. It accepts the CLI arguments
// and hands them over to the Execute function which takes it from there.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli <cmd>")
		os.Exit(1)
	}

	_, err := cmd.Execute(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
