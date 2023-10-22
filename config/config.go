package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"
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

type History struct {
	Entries []HistoryEntry `json:"enties"`
}

type HistoryEntry struct {
	Date        time.Time `json:"date"`
	SpentTotal  float64   `json:"spent_total"`
	GainedTotal float64   `json:"gained_total"`
	Expenses    Entry     `json:"expenses"`
	Incomes     Entry     `json:"incomes"`
}

type Entry struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type MonthlyRegularExpenses struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

func CreateDefaultConfigFile() error {
	err := makeDataFolder()
	if err != nil {
		return fmt.Errorf("error initialize app data: %w", err)
	}

	file, err := os.Create(CONFIG_FILE)
	if err != nil {
		return fmt.Errorf("error creating default config file: %w", err)
	}

	err = createHistoryFile()
	if err != nil {
		return fmt.Errorf("error creating history file: %w", err)
	}

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

func createHistoryFile() error {
	file, err := os.Create(HISTORY_FILE)
	if err != nil {
		return fmt.Errorf("error creating history file: %w", err)
	}

	encoded, err := json.Marshal(HistoryEntry{})
	if err != nil {
		return fmt.Errorf("error marshaling history file data on init: %w", err)
	}

	_, err = file.Write(encoded)
	if err != nil {
		return fmt.Errorf("error writing history file data on init: %w", err)
	}

	return nil
}

func makeDataFolder() error {
	if USER_HOME == "" {
		usrHome, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting user home dir: %w", err)
		}
		USER_HOME = usrHome
	}

	err := os.Mkdir(USER_HOME+"/"+INCOMER_DATA_DIR, fs.ModeDir)
	if err != nil {
		return fmt.Errorf("error can't create app data directory: %w", err)
	}

	return nil
}
