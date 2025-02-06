/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/SalmandaAK/expense-tracker/internal/config"
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
		s := config.InitiateExpenseErvice(config.FilePath)
		err := s.DeleteExpense(id)
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayMessage(fmt.Sprintln("Expense deleted successfully"))
	},
}

var id int

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVar(&id, "id", 0, "ID of expense")
	deleteCmd.MarkFlagRequired("id")
}
