package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	CONFIG_FILE  = "CONFIG.json"
	HISTORY_FILE = "HISTORY.json"

	INCOMER_DATA_DIR = "incomer_data"

	HISTORY_FILE_PATH_ARG = "history_file_path"
)

var USER_HOME string

type Config struct {
	MonthlyIncomeTotal          float64                          `json:"monthly_income"`
	IncomeMonthDaysAndAmount    map[uint8]float64                `json:"income_days_amount"`
	MonthlyExpensesTotal        float64                          `json:"monthly_expenses"`
	MonthlyRegularExpensesDates map[uint8]MonthlyRegularExpenses `json:"regular_expenses"`
	HistoryFilePath             string                           `json:"history_file_path"`
}

type MonthlyRegularExpenses struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

func CreateDefaultConfigFile() error {
	// err := makeDataFolder()
	// if err != nil {
	// 	return fmt.Errorf("error initialize app data: %w", err)
	// }

	file, err := os.Create(CONFIG_FILE)
	if err != nil {
		return fmt.Errorf("error creating default config file: %w", err)
	}

	// err = createHistoryFile()
	// if err != nil {
	// 	return fmt.Errorf("error creating history file: %w", err)
	// }

	configData := Config{
		HistoryFilePath: HISTORY_FILE,
	}

	encoded, err := json.Marshal(&configData)
	if err != nil {
		return fmt.Errorf("error marshaling default config file data: %w", err)
	}

	_, err = file.Write(encoded)
	if err != nil {
		return fmt.Errorf("error writing default config data to file: %w", err)
	}

	return nil
}
