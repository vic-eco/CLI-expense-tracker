package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

func NewAddCMD() *cobra.Command {
	var amount float64
	var description string
	
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds expenses",
		Long:  "The add command is used to add expenses to your history",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.AddExpenses(description, amount)
		},
	}

	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Amount of the expense")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of expense")

	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("description")

	return addCmd
}
