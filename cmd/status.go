/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"task_tracker/helper"

	"github.com/spf13/cobra"
)

// listDoneCmd lists only completed tasks
var listDoneCmd = &cobra.Command{
	Use:   "list-done",
	Short: "List all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := helper.LoadTasks()
		found := false
		fmt.Println("Completed Tasks:")
		for _, task := range tasks {
			if task.Status == "done" {
				fmt.Printf("[ID: %d] %s\n", task.ID, task.Name)
				found = true
			}
		}
		if !found {
			fmt.Println("No completed tasks found.")
		}
	},
}

// listInProgressCmd lists only in-progress tasks
var listInProgressCmd = &cobra.Command{
	Use:   "list-in-progress",
	Short: "List all in-progress tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := helper.LoadTasks()
		found := false
		fmt.Println("In-Progress Tasks:")
		for _, task := range tasks {
			if task.Status == "in-progress" {
				fmt.Printf("[ID: %d] %s\n", task.ID, task.Name)
				found = true
			}
		}
		if !found {
			fmt.Println("No in-progress tasks found.")
		}
	},
}

// listTodoCmd lists only tasks that are yet to be done
var listTodoCmd = &cobra.Command{
	Use:   "list-todo",
	Short: "List all tasks that are not done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := helper.LoadTasks()
		found := false
		fmt.Println("To-Do Tasks:")
		for _, task := range tasks {
			if task.Status == "todo" {
				fmt.Printf("[ID: %d] %s\n", task.ID, task.Name)
				found = true
			}
		}
		if !found {
			fmt.Println("No to-do tasks found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listDoneCmd)
	rootCmd.AddCommand(listInProgressCmd)
	rootCmd.AddCommand(listTodoCmd)
}
