package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kunalkapadia/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Successfully added %s to the task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
