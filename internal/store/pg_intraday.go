package store

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PgIntradayStore struct {
	db  *sqlx.DB
	log *logrus.Entry
}

// NewPgIntradayStore configures a PgIntradayStore with a db and returns it
func NewPgIntradayStore(db *sqlx.DB, log *logrus.Entry) *PgIntradayStore {
	return &PgIntradayStore{
		db:  db,
		log: log,
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
func (p PgIntradayStore) List(ctx context.Context, limit, offset int) ([]messages.IntradayValue, error) {
	query := sq.
		Select("ticker", "timestamp", "open", "high", "low", "close", "volume").
		From("intraday").
		OrderBy("timestamp ASC").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	q, _, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to generate SQL from query: %w", err)
	}

	query = query.RunWith(p.db)

	logrus.Info("query: ", q)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to select list from DB: %w", err)
	}
	defer rows.Close()

	list := []messages.IntradayValue{}
	for rows.Next() {
		v := messages.IntradayValue{}
		err := rows.Scan(&v.Ticker, &v.Timestamp, &v.Open, &v.High, &v.Low, &v.Close, &v.Volume)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row to value: %w", err)
		}
		list = append(list, v)
	}

	return list, nil
}
