package cmd

import (
	"fmt"
	"github.com/jkuma/gophercises/task/repository"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var cmdCompleted = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following completed tasks:")
		tasks, err := repository.ListCompletedTasks()

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for _, t := range tasks {
			fmt.Println(time.Unix(t.Time, 0).Format(time.RFC822), t.Task.Number, t.Task.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdCompleted)
}

