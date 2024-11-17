package task

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// AddTask adds a new task to the list of todos for a user.
func AddTask(taskCmd []string) (taskId int, err error) {
	if len(taskCmd) < 2 || taskCmd[1] == "" {
		err = errors.New("Task name must be provided")
		return
	}

	taskName := strings.Join(taskCmd[1:], " ")

	task := Task{
		Description: taskName,
		Status:      "todo", // default action is todo
	}

	taskId, err = task.save(nil) // persist the task data to the storage

	if err != nil {
		fmt.Fprint(os.Stderr, "Task could not be saved! Try again.")
	} else { // task was successfully saved to storage.
		fmt.Printf("Task added successfully (ID: %d)\n", taskId)
	}
	return
}
