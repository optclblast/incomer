package config

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

type StateJSON struct {
	CustomDBpath string `json:"custom_db_path"`
}
