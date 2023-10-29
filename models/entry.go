package models

import "time"

type Entry struct {
	Date    time.Time
	Title   string
	Expense float64
	Income  float64
}
