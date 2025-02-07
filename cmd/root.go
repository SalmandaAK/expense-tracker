/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SalmandaAK/expense-tracker/internal/expense/db"
	"github.com/SalmandaAK/expense-tracker/internal/expense/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile = filepath.Join(filepath.Dir(""), ".expense-tracker.yaml")

var expenseStorage = filepath.Join(filepath.Dir(""), "expense.json")
var expenseRepo = db.New(expenseStorage)
var expenseService = service.New(expenseRepo)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense-tracker is an app to track your expenses",
	Long: `
	Expense-tracker is an app to track your expenses. It provides you the ability to save your expenses,
	and give you a summary of your expenses.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetDefault("currency", "$")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".expense-tracker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".expense-tracker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
