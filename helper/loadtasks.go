package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks() []Task {
	// Check if file exists, create if not
	//This code checks if the tasks.json file exists, and if it does not exist, it creates an empty JSON file.
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		_ = os.WriteFile("tasks.json", []byte("[]"), 0644)
	}

	// Read file contents
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return []Task{} //Returns an empty task list ([]Task{}) to prevent a crash.
	}

	// Parse JSON
	//This code converts JSON data from tasks.json into a Go struct ([]Task) while handling errors properly.
	var tasks []Task                   //Declares a slice of Task structs ([]Task{}) to store the tasks.
	err = json.Unmarshal(file, &tasks) // converting json to struck
	if err != nil {
		fmt.Println("Error decoding tasks JSON:", err)
		return []Task{}
	}

	return tasks
}
