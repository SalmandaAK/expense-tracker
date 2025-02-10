/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long: `
	Delete an expense.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		s := expenseService
		err := s.DeleteExpense(id)
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayMessage("Expense deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVar(&id, "id", 0, "ID number of expense")
	deleteCmd.MarkFlagRequired("id")
}
