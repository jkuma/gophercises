package cmd

import (
	"fmt"
	"github.com/jkuma/gophercises/task/task"
	"github.com/spf13/cobra"
	"log"
)

var cmdCompleted = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following completed tasks:")
		tasks, err := task.ListCompletedTasks()

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for i, t := range tasks {
			fmt.Println(i, t.Number, t.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdCompleted)
}

