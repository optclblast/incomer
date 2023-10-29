/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package income

import (
	"context"
	"fmt"
	"incomer/models"
	"incomer/storage"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	incomeTitle   string
	incomeAmount  float64
	incomeDate    time.Time
	incomeDateStr string
)

var IncomeCmd = &cobra.Command{
	Use:   "income",
	Short: "income",
	Long:  `income long`,
	Run:   RunIncome,
}

func init() {
	IncomeCmd.Flags().StringVarP(
		&incomeTitle,
		"title",
		"t",
		"",
		"Income description",
	)
	IncomeCmd.Flags().Float64VarP(
		&incomeAmount,
		"amount",
		"a",
		0.0,
		"Income amount",
	)
	IncomeCmd.Flags().StringVarP(
		&incomeDateStr,
		"date",
		"d",
		"",
		"Date of an income [day.month.year]",
	)

	if err := IncomeCmd.MarkFlagRequired("amount"); err != nil {
		fmt.Println("amount flag is required")
		os.Exit(0)
	}
}

func RunIncome(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if incomeDateStr != "" {
		dateRegex := regexp.MustCompile(`^(\d{1,2}).(\d{1,2}).(\d{4})$`)
		if dateRegex.MatchString(incomeDateStr) {
			matches := dateRegex.FindStringSubmatch(incomeDateStr)

			day, err := strconv.Atoi(matches[1])
			if err != nil {
				fmt.Println("Invalid date format")
				os.Exit(0)
			}

			month, err := strconv.Atoi(matches[2])
			if err != nil {
				fmt.Println("Invalid date format")
				os.Exit(0)
			}

			year, err := strconv.Atoi(matches[3])
			if err != nil {
				fmt.Println("Invalid date format")
				os.Exit(0)
			}

			incomeDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		} else {
			fmt.Println("Invalid date format")
			os.Exit(0)
		}
	}

	err := storage.GlobalStorage.Init(ctx)
	if err != nil {
		fmt.Println("database error: ", err)
		os.Exit(0)
	}

	err = storage.GlobalStorage.NewIncomeEntry(ctx, models.Entry{
		Date:   incomeDate,
		Title:  incomeTitle,
		Income: incomeAmount,
	})
	if err != nil {
		fmt.Println("database error: ", err)
		os.Exit(0)
	}
}
