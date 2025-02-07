/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long: `
	Add an expense.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		s := expenseService
		id, err := s.AddExpense(description, amount)
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayMessage(fmt.Sprintf("Expense added successfully (ID: %d)\n", id))
	},
}

var description string
var amount int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&description, "description", "", "description of expense")
	addCmd.Flags().IntVar(&amount, "amount", 0, "amount of expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
