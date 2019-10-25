package cmd

import (
	"fmt"
	"github.com/jkuma/gophercises/task/task"
	"github.com/spf13/cobra"
	"strconv"
)

var cmdDo = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task on your TODO list as complete",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		k, err := strconv.Atoi(args[0])

		if err != nil {
			println(err)
		}

		t, err := task.DeleteTask(k)

		if err != nil {
			println(err)
		}

		fmt.Println("Task done:", t.Name)
	},
}

func init() {
	rootCmd.AddCommand(cmdDo)
}
