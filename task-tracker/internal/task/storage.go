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
var nextAvailableId int

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
		var fileContent TaskStorage

		err = json.Unmarshal(file, &fileContent)
		if err != nil {
			log.Fatal(err)
		}

		tempStorage = fileContent.Tasks
		nextAvailableId = fileContent.Metadata.NextId
	} else {
		nextAvailableId = 1
	}
}

// save persists the tasks to the storage.
func (task *Task) save(options map[string]bool) (taskId int, err error) {
	currentTime := time.Now()

	if options != nil && options["update"] {
		task.UpdatedAt = currentTime
	} else {
		task.Id = nextId()
		task.CreatedAt = currentTime
		task.UpdatedAt = currentTime
	}

	taskId = task.Id

	tempStorage[itoa(taskId)] = *task

	// Persist the task data and metadata
	err = persistTasks()
	if err != nil {
		return 0, err
	}

	return
}

// persistTasks serializes the current task data and metadata, and writes it
// to the storage file.
func persistTasks() error {
	// Create the task storage structure
	taskStorage := TaskStorage{
		Metadata: struct {
			NextId int `json:"nextId"`
		}{NextId: nextAvailableId}, // Using the global nextAvailableId value
		Tasks: tempStorage,
	}

	// Serialize the data
	jsonData, err := json.MarshalIndent(
		taskStorage, "", "  ",
	)
	if err != nil {
		return fmt.Errorf("persist: error serializing tasks: %w", err)
	}

	// Write to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("persist: error writing to storage file: %w", err)
	}

	return nil
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

// getById retrieves the task by its ID.
func getById(id int) (task Task, err error) {
	strId := itoa(id)
	task, found := tempStorage[strId]

	if !found {
		err = fmt.Errorf("Task with ID %d not found", id)
	}

	return
}
