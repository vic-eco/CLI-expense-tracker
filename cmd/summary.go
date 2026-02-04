package cmd

import (
	"expense-tracker/internal"
	"github.com/spf13/cobra"
)

var month int

func NewSummaryCMD() *cobra.Command {
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
