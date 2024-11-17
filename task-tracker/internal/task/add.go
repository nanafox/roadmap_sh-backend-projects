package task

import (
	"errors"
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

	taskId, err = task.save() // persist the task data to the storage
	return
}