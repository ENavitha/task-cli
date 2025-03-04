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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "markDone",
	Short: "mark a task as done  ",
	//Long:  `marking the progress tasks.`,
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
				tasks[i].Status = "done"
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

		fmt.Printf("Task %d marked as done successfully!\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
	// markDoneCmd.Flags().IntVarP(&markDoneID, "id", "i", 0, "Task ID")
	// markDoneCmd.MarkFlagRequired("id")

}
