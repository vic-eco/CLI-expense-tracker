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
			fmt.Println("Command works")
		},
	}

	return rootCmd
}
