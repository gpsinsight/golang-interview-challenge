package store_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gpsinsight/go-interview-challenge/internal/store"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func Test_Insert(t *testing.T) {
	ctx := context.TODO()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	sqlxdb := sqlx.NewDb(db, "sqlmock")

	value := messages.IntradayValue{
		Ticker:    "TEST",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Open:      100,
		High:      101,
		Low:       99,
		Close:     100.76,
		Volume:    10000,
	}

	tests := []struct {
		label string
		setup func(*testing.T)
		err   error
	}{
		{
			label: "success",
			setup: func(t *testing.T) {
				mock.ExpectExec(
					regexp.QuoteMeta("INSERT INTO intraday (ticker,timestamp,open,high,low,close,volume) VALUES ($1,$2,$3,$4,$5,$6,$7)")).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			label: "error",
			setup: func(t *testing.T) {
				mock.ExpectExec(
					regexp.QuoteMeta("INSERT INTO intraday (ticker,timestamp,open,high,low,close,volume) VALUES ($1,$2,$3,$4,$5,$6,$7)")).
					WillReturnError(fmt.Errorf("insert failed"))
			},
			err: fmt.Errorf("unable to insert row: insert failed"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.label, func(t *testing.T) {
			tt.setup(t)
			writer := store.NewPgIntradayStore(sqlxdb)
			err := writer.Insert(ctx, &value)
			if err != nil {
				require.EqualError(t, err, tt.err.Error())
			}
		})
	}
}
