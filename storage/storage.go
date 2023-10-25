package storage

import (
	"context"
	"database/sql"
	"incomer/config"
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	InitNewDatabase(ctx context.Context) error
	Connect(ctx context.Context) (*sql.DB, error)
	SetConfiguration(ctx context.Context, args *config.Config) error
	GetConfiguration(ctx context.Context) (*config.Config, error)
	GetHistory(ctx context.Context) (*History, error)
	// TODO
	// UpdateHistoryEntry(ctx context.Context) error
	// WriteHistoryEntry(ctx context.Context) error
}

type sqliteStorage struct {
	//TODO
}

type History struct {
	Entries []HistoryEntry `json:"enties"`
}

type HistoryEntry struct {
	Id          uuid.UUID `json:"id"`
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

func NewStorage(ctx context.Context) Storage {
	return &sqliteStorage{}
}

func (s *sqliteStorage) InitNewDatabase(ctx context.Context) error {
	return nil
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
