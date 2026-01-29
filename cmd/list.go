package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

func NewListCMD() *cobra.Command {
	list := &cobra.Command{
		Use:   "list",
		Short: "Lists all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.ListExpenses()
		},
	}

	return list
}
