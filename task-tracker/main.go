package main

import (
	"fmt"
	"os"

	"github.com/nanafox/task-tracker/cmd"
)

// main provides the entry to the task-cli program. It accepts the CLI arguments
// and hands them over to the Execute function which takes it from there.
func main() {
	_, err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
