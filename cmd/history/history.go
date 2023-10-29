package history

import (
	"context"
	"fmt"
	"incomer/storage"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	read  bool
	clear bool

	fromS string
	toS   string

	from time.Time
	to   time.Time
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "income",
	Long:  `income long`,
	Run:   RunHistory,
}

func init() {
	HistoryCmd.Flags().BoolVarP(
		&read,
		"read",
		"r",
		true,
		"Read history",
	)
	HistoryCmd.Flags().BoolVarP(
		&clear,
		"clear",
		"c",
		false,
		"Clears history",
	)
	HistoryCmd.Flags().StringVarP(
		&fromS,
		"from",
		"f",
		"",
		"Date from [day.month.year]",
	)
	HistoryCmd.Flags().StringVarP(
		&toS,
		"to",
		"t",
		"",
		"Date to [day.month.year]",
	)
}

func RunHistory(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if fromS != "" {
		dateRegex := regexp.MustCompile(`^(\d{1,2}).(\d{1,2}).(\d{4})$`)
		if dateRegex.MatchString(fromS) {
			matches := dateRegex.FindStringSubmatch(fromS)

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

			from = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		} else {
			fmt.Println("Invalid date format")
			os.Exit(0)
		}
	}

	if toS != "" {
		dateRegex := regexp.MustCompile(`^(\d{1,2}).(\d{1,2}).(\d{4})$`)
		if dateRegex.MatchString(toS) {
			matches := dateRegex.FindStringSubmatch(toS)

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

			to = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		} else {
			fmt.Println("Invalid date format")
			os.Exit(0)
		}
	}

	if read {
		err := storage.GlobalStorage.Init(ctx)
		if err != nil {
			fmt.Println("database error: ", err)
			os.Exit(0)
		}

		historyEntries, err := storage.GlobalStorage.GetHistory(ctx, from, to)
		if err != nil {
			fmt.Println("database error: ", err)
			os.Exit(0)
		}

		for _, e := range historyEntries.Entries {
			var amount float64
			if e.Income == 0 {
				amount = e.Expense
			} else {
				amount = e.Income
			}

			fmt.Printf(
				"[%v-%v-%v] | %s | %v\n",
				e.Date.Day(),
				e.Date.Month(),
				e.Date.Year(),
				e.Title,
				amount,
			)
		}
	}
}
