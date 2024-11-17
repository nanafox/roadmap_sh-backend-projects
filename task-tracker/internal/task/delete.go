package task

import (
	"errors"
	"fmt"
	"strconv"
)

// DeleteTask removes a task by its ID.
//
// Parameters:
//
//	cmd []string: A slice of strings containing the command. Expected format is "task-cli delete <taskId>".
//
// Returns:
//
//	taskId (int): The ID of the task that was deleted.
//	err (error): An error if the task cannot be found or if the input is invalid.
//
// Example:
//
//	cmd := []string{"delete", "2"}
//	taskId, err := DeleteTask(cmd)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Printf("Task with ID %d deleted successfully!\n", taskId)
//	}
func DeleteTask(cmd []string) (taskId int, err error) {
	if len(cmd) != 2 {
		return 0, errors.New("delete: invalid syntax. Usage: task-cli delete <taskId>")
	}

	taskId, err = strconv.Atoi(cmd[1])
	if err != nil {
		return 0, fmt.Errorf("delete: invalid taskId '%s'", cmd[1])
	}

	_, exists := tempStorage[itoa(taskId)]
	if !exists {
		return 0, fmt.Errorf("delete: task with ID %d not found", taskId)
	}

	delete(tempStorage, itoa(taskId))

	// Persist the task data and metadata after deletion
	err = persistTasks()
	if err != nil {
		return 0, err
	}

	fmt.Printf("Task with ID %d deleted successfully!\n", taskId)
	return taskId, nil
}
