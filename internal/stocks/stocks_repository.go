package stocks

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	logger *logrus.Entry
	db     *sqlx.DB
}

func NewRepository(logger *logrus.Entry, db *sqlx.DB) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

type IntradayValue struct {
	Ticker    string
	Timestamp string
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    int64
}

func (r *Repository) InsertIntradayValue(value *IntradayValue) error {
	// TODO: perform db insert
	return nil
}
