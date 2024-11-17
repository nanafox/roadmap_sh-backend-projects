package task

import (
	"fmt"
	"strconv"
)

// MarkTask updates the status of an existing task.
//
// Usage:
//
//	MarkTask([]string{"mark-in-progress", "<taskId>"})
//	MarkTask([]string{"mark-done", "<taskId>"})
//
// Parameters:
//
//	cmd []string: A slice of strings representing the command arguments.
//	              The first element specifies the action ("mark-in-progress" or "mark-done"),
//	              and the second element is the task ID.
//
// Returns:
//
//	taskId (int): The ID of the updated task if the operation is successful.
//	err (error): An error if the input is invalid, the task is not found,
//	             or there is an issue saving the task.
//
// Behavior:
//  1. Validates the input `cmd` to ensure it contains exactly 2 elements.
//  2. Parses the task ID from the second element of `cmd`.
//  3. Retrieves the task by ID.
//  4. Updates the task's status to "in-progress" or "done" based on the action in `cmd[0]`.
//  5. Saves the updated task to persistent storage.
//
// Errors:
//   - Returns an error if the `cmd` slice length is not 2.
//   - Returns an error if the task ID cannot be parsed as an integer.
//   - Returns an error if the task cannot be retrieved by the provided ID.
//   - Returns an error if the specified action is invalid.
//   - Returns an error if saving the updated task fails.
//
// Example:
//
//	cmd := []string{"mark-in-progress", "1"}
//	taskId, err := MarkTask(cmd)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Task status updated! Task ID:", taskId)
//	}
func MarkTask(cmd []string) (taskId int, err error) {
	const usage = "task-cli mark-{in-progress | done} <taskId>"

	if len(cmd) != 2 {
		return 0, fmt.Errorf("invalid input. Expected format: %s", usage)
	}

	// Extract action and parse task ID
	markAction, taskIdStr := cmd[0], cmd[1]
	id, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return 0, fmt.Errorf("invalid taskId '%s': %w", taskIdStr, err)
	}

	task, err := getById(id)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve task: %w", err)
	}

	// Update the task's status based on the action
	switch markAction {
	case "mark-in-progress":
		task.Status = "in-progress"
	case "mark-done":
		task.Status = "done"
	default:
		return 0, fmt.Errorf(
			"invalid action '%s'. Expected 'mark-in-progress' or 'mark-done'",
			markAction,
		)
	}

	// Save the updated task
	return task.markAndSave()
}

// markAndSave updates the task in persistent storage.
//
// Behavior:
//  1. Saves the task using the `save` method with an update option.
//  2. Prints a success message if the save is successful.
//
// Returns:
//
//	taskId (int): The ID of the saved task.
//	err (error): An error if saving the task fails.
//
// Example:
//
//	task := &Task{ID: 1, Status: "in-progress"}
//	taskId, err := task.markAndSave()
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Task status updated successfully! Task ID:", taskId)
//	}
func (task *Task) markAndSave() (taskId int, err error) {
	taskId, err = task.save(map[string]bool{"update": true})
	if err != nil {
		return 0, fmt.Errorf("failed to save task: %w", err)
	}

	fmt.Println("Task status updated successfully")
	return taskId, nil
}
