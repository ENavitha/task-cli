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

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-In-Progress",
	Short: "tasks that are in progress",

	Run: func(cmd *cobra.Command, args []string) {

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		tasks := helper.LoadTasks()

		// Find the task and update status
		found := false
		for i, task := range tasks {
			if task.ID == taskID {
				tasks[i].Status = "in-progress"
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Task not found with ID:", taskID)
			return
		}

		err = helper.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error updating task status:", err)
			return
		}

		fmt.Printf("Task %d marked as in-progress successfully!\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// markInProgressCmd.Flags().IntVarP(&markInProgressID, "id", "i", 0, "mark Id to progress ")
	// markInProgressCmd.MarkFlagRequired("id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
