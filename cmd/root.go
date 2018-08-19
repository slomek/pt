package cmd

import (
	"github.com/spf13/cobra"

	"github.com/slomek/pt/cmd/mine"
)

func init() {
	Command.AddCommand(mine.Command)

	Command.PersistentFlags().StringP("project-id", "p", "", "project ID")
	Command.PersistentFlags().StringP("username", "u", "", "current user's username")
}

// Command is a root command of the application, it deoesn't do anything on its own.
var Command = &cobra.Command{
	Use:   "pt",
	Short: "CLI client for Pivotal Tracker",
	Long:  "Command line to make Pivotal Tracker interaction faster",
}
