package cmd

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/nanafox/task-tracker/internal/task"
)

// Execute processes a command and delegates it to the appropriate action handler.
//
// Parameters:
//
//	cmd []string: A slice of strings representing the user command.
//	              The first element specifies the action, and subsequent elements are arguments.
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
func Execute(cmd []string) (taskId int, err error) {
	if len(cmd) == 0 {
		return 0, errors.New("no command provided")
	}

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
		"invalid action '%s'. Valid actions are: %v",
		action, validActions,
	)
}
