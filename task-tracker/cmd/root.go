package cmd

import (
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/nanafox/task-tracker/internal/task"
)

// Execute processes a command and delegates it to the appropriate action handler.
//
// Returns:
//
//	taskId (int): The ID of the task processed by the handler, if applicable.
//	err (error): An error if the command is invalid or if the action handler fails.
//
// Behavior:
//  1. Extracts the action from the first element of `cmd`.
//  2. Delegates the command to the appropriate handler based on the action.
//  3. Returns an error if the action is invalid or if the handler fails.
//
// Example:
//
//	cmd := []string{"add", "New Task"}
//	taskId, err := Execute(cmd)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Task successfully processed! Task ID:", taskId)
//	}
func Execute() (taskId int, err error) {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		displayHelp()
		os.Exit(1)
	}

	cmd := os.Args[1:]
	action := cmd[0]

	// Map of actions to their corresponding handlers
	actionHandlers := map[string]func([]string) (int, error){
		"add":              task.AddTask,
		"list":             task.ListAll,
		"update":           task.Update,
		"mark-in-progress": task.MarkTask,
		"mark-done":        task.MarkTask,
		"delete":           task.DeleteTask,
	}

	// Find and execute the corresponding handler
	if handler, exists := actionHandlers[action]; exists {
		return handler(cmd)
	}

	// Handle invalid action
	validActions := slices.Collect(maps.Keys(actionHandlers))

	return 0, fmt.Errorf(
		"invalid action '%s'. Valid actions are: %v.\n\nUse '%s help' to get help",
		action, validActions, os.Args[0],
	)
}

func displayHelp() {
	fmt.Println(`
TASK TRACKER CLI
================
A simple command-line application to manage your tasks efficiently.

USAGE:
  task-cli [COMMAND] [OPTIONS]

COMMANDS:
  add               Add a new task to the tracker.
                    Example: task-cli add "Buy groceries"

  list              List all tasks, or specific tasks.
                    Example: 
                      - task-cli list
                      - task-cli list todo # lists tasks that have started yet.
                      - task-cli list in-progress # lists tasks in progress
                      - task-cli list done # lists completed tasks.

  update            Update an existing task.
                    Example: task-cli update 1 "Buy groceries and fruits"

  mark-done         Marks a task as completed.
                    Example: task mark-done 1

  mark-in-progress  Marks a task to be in progress.
                    Example: task mark-in-progress 4

  delete            Delete a task by its ID.
                    Example: task-cli delete 1

  help              Display this help message.
                    Example: task-cli help

EXAMPLES:
  1. Add a new task:
     task-cli add "Complete homework"
     # Output: Task created successfully (ID: 1)

  2. List all tasks:
     task-cli list

  3. Update an existing task:
     task-cli update 3 "Submit project report"

  4. Delete a task:
     task-cli delete 5

  5. Get help:
     task-cli help

NOTES:
- Task IDs are generated automatically and can be used with 'update' or 'delete'.
- Use descriptive names for tasks for better clarity.

CONTACT:
For issues or feedback, please visit the GitHub repository:
https://github.com/nanafox/task-tracker

Happy task managing!`)
}
