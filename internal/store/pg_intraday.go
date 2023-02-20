package store

import (
	"context"

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
func (p PgIntradayStore) Insert(ctx context.Context, val messages.IntradayValue) error {
	return nil
}

// List returns a collection of IntradayValue from the DB
func (p PgIntradayStore) List(ctx context.Context) ([]messages.IntradayValue, error) {
	return nil, nil
}
