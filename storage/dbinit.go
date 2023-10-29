package storage

const (
	qinit_config = `CREATE TABLE IF NOT EXISTS config (
		'monthly_income' REAL DEFAULT 0,
		'monthly_expenses' REAL DEFAULT 0,
		'history_file_path' TEXT
	);`
	qinit_income_days = `CREATE TABLE IF NOT EXISTS income_days_amount (
		'datahash' BLOB NOT NULL,
		'month_day' INTEGER NOT NULL,
		'amount' REAL
	);`
	qinit_dated_regular_expenses = `CREATE TABLE IF NOT EXISTS dated_regular_expenses (
		'datahash' BLOB NOT NULL,
		'month_day' INTEGER NOT NULL,
		'data_tag' BLOB
	);`
	qinit_regular_expenses_data = `CREATE TABLE IF NOT EXISTS regular_expenses_data (
		'datahash' BLOB NOT NULL,
		'title' INTEGER NOT NULL,
		'amount' REAL DEFAULT 0
	);`
	qinit_history_entries = `CREATE TABLE IF NOT EXISTS history_entries (
		'date' DATETIME,
		'title' TEXT,
		'expense' REAL DEFAULT 0,
		'income' REAL DEFAULT 0
	);`
)
