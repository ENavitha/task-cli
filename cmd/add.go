/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"task_tracker/helper"

	"github.com/spf13/cobra"
)

var TaskName string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := helper.LoadTasks()
		TaskName = args[0]
		// Create a new task with a unique ID
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}

		newTask := helper.Task{
			ID:     maxID + 1,
			Name:   TaskName,
			Status: "todo",
		}
		tasks = append(tasks, newTask)
		helper.SaveTasks(tasks)

		// Save tasks to file
		err := helper.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("Task added successfully:", TaskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	//addCmd.Flags().StringVarP(&helper.TaskName, "name", "n", "", "Task name")
	//addCmd.MarkFlagRequired("name")
}

// Load tasks from tasks.json, create file if missing
