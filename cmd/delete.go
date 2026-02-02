package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

var id int

func NewDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes an expense",
		Long:  "The delete commands is used to delete an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.DeleteExpense(id)
		},
	}

	deleteCmd.Flags().IntVar(&id, "id", 0, "id of the expense to be deleted")
	deleteCmd.MarkFlagRequired("id")

	return deleteCmd
}
