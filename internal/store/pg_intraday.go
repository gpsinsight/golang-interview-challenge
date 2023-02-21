package store

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/jmoiron/sqlx"
)

type PgIntradayStore struct {
	db *sqlx.DB
}

// NewPgIntradayStore configures a PgIntradayStore with a db and returns it
func NewPgIntradayStore(db *sqlx.DB) *PgIntradayStore {
	return &PgIntradayStore{
		db: db,
	}
}

// Insert inserts a new IntradayValue row into the DB
func (p PgIntradayStore) Insert(ctx context.Context, val *messages.IntradayValue) error {
	ts, err := time.Parse("2006-01-02 15:04:05", val.Timestamp)
	if err != nil {
		return fmt.Errorf("unable to parse value timestamp: %w", err)
	}

	query := sq.
		Insert("intraday").
		Columns("ticker", "timestamp", "open", "high", "low", "close", "volume").
		Values(val.Ticker, ts, val.Open, val.High, val.Low, val.Close, val.Volume).
		RunWith(p.db).
		PlaceholderFormat(sq.Dollar)

	_, err = query.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

// List returns a collection of IntradayValue from the DB
func (p PgIntradayStore) List(ctx context.Context) ([]messages.IntradayValue, error) {
	return nil, nil
}
