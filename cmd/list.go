/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/SalmandaAK/expense-tracker/internal/config"
	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Long: `
	List all expenses.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		s := config.InitiateExpenseErvice(config.FilePath)
		expenses, err := s.FindAllExpenses()
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayExpenseList(expenses)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
