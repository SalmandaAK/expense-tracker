package cmd

import (
	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show summary of all expenses, or show summary by month",
	Long: `
	Show summary of expenses. Can also be used to show summary of expense within a month.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		s := expenseService
		if month != 0 {
			summary, err := s.SummaryExpensesByMonth(month)
			if err != nil {
				view.DisplayError(err)
				return
			}
			view.DisplaySummary(summary, month)
		} else {
			summary, err := s.SummaryExpenses()
			if err != nil {
				view.DisplayError(err)
				return
			}
			view.DisplaySummary(summary, month)
		}
	},
}

var month int

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVar(&month, "month", 0, "month number where the expenses happened")
}
