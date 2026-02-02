package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "Expense Tracker is a CLI tool for expenses",
		Long:  "Expense Tracker is a CLI tool that allows for personal expenses tracking",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Expense Tracker")
		},
	}

	rootCmd.AddCommand(NewListCMD())
	rootCmd.AddCommand(NewAddCMD())
	rootCmd.AddCommand(NewDeleteCmd())

	return rootCmd
}
