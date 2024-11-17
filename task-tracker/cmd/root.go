package cmd

import (
	"fmt"

	"github.com/nanafox/task-tracker/internal/task"
)

// Execute handles the command received on the CLI
//
// It parses it and hands control to the appropriate function to complete the
// action.
func Execute(cmd []string) (taskId int, err error) {
	action := cmd[0]

	switch action {
	case "add":
		return task.AddTask(cmd)
	case "list":
		return task.ListAll(cmd)
	case "update":
		return task.Update(cmd)
	}

	return 0, fmt.Errorf("hanlde_cmd: %s is not a valid action", action)
}
