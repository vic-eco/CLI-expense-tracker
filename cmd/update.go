package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

func NewUpdateCMD() *cobra.Command {
	var amount float64
	var description string
	var id int

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Updates an expense",
		Long:  "The update command is used to update an expense by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.UpdateExpense(id, description, amount)
		},
	}

	updateCmd.Flags().IntVar(&id, "id", 0, "id of expense to update")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "updated description")
	updateCmd.Flags().Float64VarP(&amount, "amount", "a", -1, "updated amount")

	updateCmd.MarkFlagRequired("id")

	return updateCmd
}
