package task

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"time"
)

var tempStorage map[string]Task

const filename string = "tasks.json"

// init initializes the storage for storing the tasks.
func init() {
	createFile(filename)

	tempStorage = make(map[string]Task)

	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	if len(file) != 0 {
		err = json.Unmarshal(file, &tempStorage)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// save persists the tasks to the storage.
func (task *Task) save(options map[string]bool) (taskId int, err error) {
	currentTime := time.Now()

	if options != nil && options["update"] {
		task.UpdatedAt = time.Now()
	} else {
		task.Id = nextId()
		task.CreatedAt = currentTime
		task.UpdatedAt = currentTime
	}

	taskId = task.Id

	tempStorage[itoa(taskId)] = *task

	jsonData, err := json.MarshalIndent(tempStorage, "", "  ")

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return 0, err
	}

	return
}

// getAll returns all the tasks in the storage.
//
// When the status is a non-empty string, it is used to filter the tasks returned
// based on its status. The list of accepted task statuses are as follows:
//
// todo => For tasks not yet started.
// in-progress => For tasks that has been started but not completed.
// done => For completed tasks.
func getAll(status string) (tasks []Task, err error) {
	allTasks := slices.Collect(maps.Values(tempStorage))

	if status == "" {
		tasks = allTasks
	} else {
		tasks, err = getByStatus(status)
	}
	return
}

// getByStatus returns all tasks that has a specific status.
func getByStatus(status string) (tasks []Task, err error) {
	allTasks := slices.Collect(maps.Values(tempStorage))

	if !slices.Contains(AllowedStatuses, status) {
		err = fmt.Errorf("get_by_status: %s is not a valid status", status)
	}
	for _, task := range allTasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return
}

func getById(id int) (task Task, err error) {
	strId := itoa(id)
	task, found := tempStorage[strId]

	if !found {
		err = fmt.Errorf("Task with ID %d not found", id)
	}

	return
}
