package cmd

import (
	"fmt"
	"github.com/jkuma/gophercises/task/repository"
	"github.com/spf13/cobra"
	"log"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following tasks:")
		tasks, err := repository.ListTasks()

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for _, t := range tasks {
			fmt.Println(t.Number, ".", t.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdList)
}
