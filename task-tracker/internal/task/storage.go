package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
func (task *Task) save() (taskId int, err error) {
	currentTime := time.Now()

	task.Id = nextId()
	task.CreatedAt = currentTime
	task.UpdatedAt = currentTime

	taskId = task.Id

	tempStorage[itoa(taskId)] = *task

	jsonData, err := json.MarshalIndent(
		tempStorage,
		"",
		"  ",
	)

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		return 0, err
	}

	fmt.Println("Saved successfully")
	return
}
