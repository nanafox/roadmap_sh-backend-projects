package task

import (
	"fmt"
	"strconv"
)

// Update modifies the description of an existing task.
//
// Usage:
//
//	Update([]string{"update", "<taskId>", "<new description>"})
//
// Parameters:
//
//	cmd []string: A slice of strings representing the command arguments.
//	              The expected format is: ["update", "<taskId>", "<new description>"].
//
// Returns:
//
//	taskId (int): The ID of the updated task if the operation is successful.
//	err (error): An error if the input is invalid, the task is not found,
//	             or there is an issue saving the task.
//
// Behavior:
//  1. Validates the input `cmd` to ensure it contains exactly 3 elements.
//  2. Parses the second argument as an integer (`taskId`).
//  3. Attempts to retrieve the task using `taskId`.
//  4. Updates the task's description with the third argument in `cmd`.
//  5. Saves the updated task to persistent storage.
//
// Errors:
//   - Returns an error if the `cmd` slice length is not 3.
//   - Returns an error if `taskId` cannot be parsed as an integer.
//   - Returns an error if the task cannot be retrieved by the provided ID.
//   - Returns an error if saving the updated task fails.
//
// Example:
//
//	cmd := []string{"update", "1", "New Task Description"}
//	taskId, err := Update(cmd)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Task updated successfully! Task ID:", taskId)
//	}
func Update(cmd []string) (taskId int, err error) {
	const expectedFormat = "task-cli update <taskId> <new description>"

	// Validate input length
	if len(cmd) != 3 {
		return 0, fmt.Errorf("update: invalid input.\n Expected format: %s", expectedFormat)
	}

	// Parse taskId
	id, parseErr := strconv.Atoi(cmd[1])
	if parseErr != nil {
		return 0, fmt.Errorf("update: invalid taskId '%s'", cmd[1])
	}

	// Fetch the task
	task, err := getById(id)
	if err != nil {
		return 0, fmt.Errorf("update: %w", err)
	}

	// Update and save the task
	task.Description = cmd[2]
	if _, saveErr := task.save(map[string]bool{"update": true}); saveErr != nil {
		return 0, fmt.Errorf("update: failed to save task: %w", saveErr)
	}

	fmt.Println("Task updated successfully")

	return id, nil
}
