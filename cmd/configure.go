/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/SalmandaAK/expense-tracker/internal/view"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure currency shown",
	Long: `
	Configure lets you type the currency you are using. Example usage:
	./expense-tracker configure --currency IDR
	The default is $
	`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("currency", currency)
		err := viper.WriteConfig()
		if err != nil {
			view.DisplayError(err)
			return
		}
		view.DisplayMessage(fmt.Sprintf("Currency has been changed into %s successfully", currency))
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVar(&currency, "currency", "", "Change currency")
	configureCmd.MarkFlagRequired("currency")
}
