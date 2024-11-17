package main

import (
	"fmt"
	"os"

	"github.com/nanafox/task-tracker/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli <cmd>")
		os.Exit(1)
	}

	taskId, err := cmd.Execute(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Task created with ID: %d", taskId)
}
