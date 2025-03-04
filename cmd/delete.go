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

//var deleteID int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task ",

	Run: func(cmd *cobra.Command, args []string) {

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		tasks := helper.LoadTasks()

		// Find and remove the task
		newTasks := []helper.Task{}
		found := false

		for _, task := range tasks {
			if task.ID == taskID {
				found = true
			} else {
				newTasks = append(newTasks, task)
			}
		}

		if !found {
			fmt.Println("Task not found with ID:", taskID)
			return
		}

		// Save the updated task list
		err = helper.SaveTasks(newTasks)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	// deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID of the task to delete")
	// deleteCmd.MarkFlagRequired("id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
