package cmd

import (
	"fmt"
	"github.com/jkuma/gophercises/task/repository"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var cmdPrint = &cobra.Command{
	Use:   "add [task to add]",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := strings.Join(args, " ")
		err := repository.AddTask(t)

		if err != nil {
			log.Fatalf("The task %v cannot be added", t)
		}

		fmt.Println("Task added: " + t)
	},
}

func init() {
	rootCmd.AddCommand(cmdPrint)
}
