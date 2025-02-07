/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Long: `
	List all expenses.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		s := expenseService
		expenses, err := s.FindAllExpenses()
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayExpenseList(expenses, viper.GetString("currency"))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
