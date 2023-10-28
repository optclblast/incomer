package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"incomer/config"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/mattn/go-sqlite3"
)

const (
	DRIVER = "sqlite3"
)

type Storage interface {
	Init(ctx context.Context) error
	Connect(ctx context.Context) (*sql.DB, error)
	SetConfiguration(ctx context.Context, args *config.Config) error
	GetConfiguration(ctx context.Context) (*config.Config, error)
	GetHistory(ctx context.Context) (*History, error)
	// TODO
	// UpdateHistoryEntry(ctx context.Context) error
	// WriteHistoryEntry(ctx context.Context) error
}

type sqliteStorage struct {
	db *sql.DB
}

type History struct {
	Entries []HistoryEntry `json:"enties"`
}

type HistoryEntry struct {
	Id          uuid.UUID `json:"id"`
	Date        time.Time `json:"date"`
	SpentTotal  float64   `json:"spent_total"`
	GainedTotal float64   `json:"gained_total"`
	Expenses    []Entry   `json:"expenses"`
	Incomes     []Entry   `json:"incomes"`
}

type Entry struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

func NewStorage(ctx context.Context, dbFilePath string) (Storage, error) {
	if !databaseExists(dbFilePath) {
		_, err := os.Create(dbFilePath)
		if err != nil {
			return nil, fmt.Errorf("error create database file: %w", err)
		}
	}

	db, err := sql.Open(DRIVER, dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("error open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error open database: %w", err)
	}

	return &sqliteStorage{db: db}, nil
}

func (s *sqliteStorage) Connect(ctx context.Context) (*sql.DB, error) {
	return nil, nil
}

func (s *sqliteStorage) SetConfiguration(ctx context.Context, args *config.Config) error {
	return nil
}

func (s *sqliteStorage) GetConfiguration(ctx context.Context) (*config.Config, error) {
	return nil, nil
}

func (s *sqliteStorage) GetHistory(ctx context.Context) (*History, error) {
	return nil, nil
}

// True if database file exists
func databaseExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (s *sqliteStorage) Init(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return fmt.Errorf("error init transaction: %w", err)
	}

	var queries []string = []string{
		qinit_config,
		qinit_income_days,
		qinit_dated_regular_expenses,
		qinit_regular_expenses_data,
		qinit_history_entries,
		qinit_entries,
	}

	txContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return handleExecTx(txContext, tx, queries)
}

// Helps to handle exec statements transaction
func handleExecTx(ctx context.Context, tx *sql.Tx, queries []string) error {
	var err error

	for _, q := range queries {
		var stmt *sql.Stmt = new(sql.Stmt)
		stmt, err = tx.PrepareContext(ctx, q)
		if err != nil {
			err = fmt.Errorf("error preparing database initialization: %w", err)
			break
		}

		_, err = stmt.ExecContext(ctx)
		if err != nil {
			err = fmt.Errorf("error executing one of tranactions statements: %w", err)
			break
		}
	}

	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return sqlite3.ErrAbortRollback
		}
	}

	return tx.Commit()
}
