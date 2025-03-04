/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"task_tracker/helper"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks ",

	Run: func(cmd *cobra.Command, args []string) {
		tasks := helper.LoadTasks()

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println("Listing all tasks:")
		for _, task := range tasks {
			fmt.Printf(" %d.%s\n", task.ID, task.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
