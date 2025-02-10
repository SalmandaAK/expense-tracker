/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense",
	Long: `
	Update description and/or amount of an expense by ID
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := expenseService.UpdateExpense(id, amount, description)
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayMessage(fmt.Sprintf("Expense updated successfully (ID: %v)", id))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	updateCmd.Flags().IntVar(&id, "id", 0, "ID of expense to be updated")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().IntVar(&amount, "amount", 0, "amount of the new espense")
	updateCmd.Flags().StringVar(&description, "description", "", "description of the new expense")
	updateCmd.MarkFlagsOneRequired("amount", "description")
}
