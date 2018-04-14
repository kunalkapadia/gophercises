package cmd

import (
	"fmt"

	"os"
	"strconv"

	"github.com/kunalkapadia/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.GetTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("%d is <=0 or more than len %d\n", id, len(tasks))
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to delete task %d, %v\n", id, err)
			} else {
				fmt.Printf("Mark task %d as completed\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
