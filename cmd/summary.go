package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

func NewSummaryCMD() *cobra.Command {
	var month int

	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "Summarizes your expenses",
		Long:  "Summarize your expenses (all/monthly)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.SummarizeExpenses(month)
		},
	}

	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Month to filter by")

	return summaryCmd
}
