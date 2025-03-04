package helper

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// save the tasks in json file
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ") // struct to json
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
