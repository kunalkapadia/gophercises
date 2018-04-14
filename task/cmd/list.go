// Copyright © 2018 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"os"

	"github.com/kunalkapadia/gophercises/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetTasks()
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete. Why not take a vacation? 🏖")
			return
		}
		fmt.Println("You have following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
