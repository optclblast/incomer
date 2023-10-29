/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package income

import (
	"fmt"

	"github.com/spf13/cobra"
)

// incomesCmd represents the incomes command
var IncomesCmd = &cobra.Command{
	Use:   "incomes",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incomes called")
	},
}

func init() {

}

func RunIncomes(cmd *cobra.Command, args []string) {
	//storage.GlobalStorage.SetConfiguration()
}
