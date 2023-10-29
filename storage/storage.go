package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"incomer/config"
	"incomer/models"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER   = "sqlite3"
	MAXLIMIT = 500
)

var GlobalStorage Storage

type Storage interface {
	Init(ctx context.Context) error
	Connect(ctx context.Context) (*sql.DB, error)
	SetConfiguration(ctx context.Context, args *config.Config) error
	GetConfiguration(ctx context.Context) (*config.Config, error)
	GetHistory(ctx context.Context, from time.Time, to time.Time) (*History, error)
	// TODO
	NewIncomeEntry(ctx context.Context, entry models.Entry) error
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
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Expense float64   `json:"expense"`
	Income  float64   `json:"income"`
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

func (s *sqliteStorage) GetHistory(ctx context.Context, from time.Time, to time.Time) (*History, error) {
	var entries []HistoryEntry = make([]HistoryEntry, 0)
	q := "SELECT * FROM history_entries WHERE date >= ? AND date <= ? date LIMIT ?"

	rows, err := s.db.QueryContext(ctx, q, from, to, MAXLIMIT)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error select from history: %w", err)
	}

	if err = rows.Scan(entries); err != nil {
		return nil, fmt.Errorf("error scan rows: %w", err)
	}

	return &History{Entries: entries}, nil
}

func (s *sqliteStorage) NewIncomeEntry(ctx context.Context, entry models.Entry) error {
	if entry.Date.IsZero() {
		entry.Date = time.Now()
	}

	q := "INSERT INTO history_entries (date, title, income) VALUES (?, ?, ?)"

	_, err := s.db.ExecContext(ctx, q, entry.Date, entry.Title, entry.Income)
	return err
}

// True if database file exists
func databaseExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (s *sqliteStorage) Init(ctx context.Context) error {
	var queries []string = []string{
		qinit_config,
		qinit_income_days,
		qinit_dated_regular_expenses,
		qinit_regular_expenses_data,
		qinit_history_entries,
	}

	for _, q := range queries {
		err := func() error {
			qCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			_, err := s.db.ExecContext(qCtx, q)
			if err != nil {
				return fmt.Errorf("error executing database initialization query [ %s ]: %w", q, err)
			}

			return nil
		}()

		if err != nil {
			return fmt.Errorf("error initializing datanase: %w", err)
		}
	}
	return nil
}
