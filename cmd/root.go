/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	confcmd "incomer/cmd/config"
	"incomer/cmd/expenses"
	"incomer/cmd/income"
	"incomer/config"
	"incomer/storage"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "incomer",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	addSubcommands()

}

func addSubcommands() {
	rootCmd.AddCommand(income.IncomeCmd)
	rootCmd.AddCommand(income.IncomesCmd)
	rootCmd.AddCommand(expenses.ExpenseCmd)
	rootCmd.AddCommand(expenses.ExpensesCmd)
	rootCmd.AddCommand(confcmd.ConfigCmd)
}

func initConfig() {
	var databasePath string

	statefiledir, err := os.Getwd()
	if err != nil {
		panic("error can't parse statfile dir path: " + err.Error())
	}

	stateFileData, err := os.ReadFile(statefiledir + "/state.json")
	if err != nil {
		panic("error can't read from statfile.json: " + err.Error())
	}

	var stateFile config.StateJSON
	err = json.Unmarshal(stateFileData, &stateFile)
	if err != nil {
		panic("error can't parde from statfile.json data: " + err.Error())
	}

	if stateFile.CustomDBpath == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			panic("error can't parse user homedir path: " + err.Error())
		}

		databasePath = fmt.Sprintf("%s/incomer.db", homedir)
	} else {
		databasePath = stateFile.CustomDBpath
	}

	if _, err := os.Stat(databasePath); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(databasePath)
		if err != nil {
			panic("error can't create new database: " + err.Error())
		}
	}

	storage.GlobalStorage, err = storage.NewStorage(context.Background(), databasePath)
	if err != nil {
		panic(err)
	}

	err = storage.GlobalStorage.Init(context.Background())
	if err != nil {
		panic(err)
	}
}
