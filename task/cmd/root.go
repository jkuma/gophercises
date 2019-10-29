package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "task",
		Short: "task is a CLI for managing your TODOs.",
		Long: `Task is a complete task manager for your daily tasks. 
Please type help to print further information about its usage`,
		Version: "1.0",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
