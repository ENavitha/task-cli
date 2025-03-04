/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task_tracker/helper"

	"github.com/spf13/cobra"
)

// var updateID int
// var newTaskName string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task (todo, in-progress, done)",

	Run: func(cmd *cobra.Command, args []string) {

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}
		newTaskName := args[1]

		tasks := helper.LoadTasks()

		// Find and update the task
		updated := false
		for i, task := range tasks {
			if task.ID == taskID {
				tasks[i].Name = newTaskName
				updated = true
				break
			}
		}
		if !updated {
			fmt.Println("Task not found with ID:", taskID)
			return
		}
		err = helper.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}

		fmt.Println("Task updated successfully!")

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	// updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the task to update")
	// updateCmd.Flags().StringVarP(&newTaskName, "name", "n", "", "New name for the task")

	// updateCmd.MarkFlagRequired("id")
	// updateCmd.MarkFlagRequired("name")

}
