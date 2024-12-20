package task

import "time"

var AllowedStatuses = []string{"todo", "in-progress", "done"}

// Task is the struct for managing tasks
type Task struct {
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Id          int       `json:"id"`
}

// TaskStorage struct
type TaskStorage struct {
	Tasks    map[string]Task `json:"tasks"`
	Metadata struct {
		NextId int `json:"nextId"`
	} `json:"metadata"`
}
