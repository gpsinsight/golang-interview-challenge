package server_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/gpsinsight/go-interview-challenge/internal/server"
	"github.com/gpsinsight/go-interview-challenge/internal/store"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testLogger() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log := logger.WithField("environment", "testing")

	return log
}

func Test_GetIntradayHandler(t *testing.T) {
	router := mux.NewRouter()

	req, err := http.NewRequest("GET", "/api/v1/intraday?page=2", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	log := testLogger()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	sqlxdb := sqlx.NewDb(db, "sqlmock")
	t.Run("valid row return = valid response", func(t *testing.T) {
		mockRows := sqlmock.NewRows([]string{"ticker", "timestamp", "open", "high", "low", "close", "volume"}).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000).
			AddRow("FOO", "2023-02-21 15:00:00", 100, 101, 99, 100, 1000)

		mock.ExpectQuery(
			regexp.QuoteMeta("SELECT ticker, timestamp, open, high, low, close, volume FROM intraday ORDER BY timestamp ASC LIMIT 20 OFFSET 20")).
			WillReturnRows(mockRows)

		s := store.NewPgIntradayStore(sqlxdb, log)

		router.HandleFunc("/api/v1/intraday", server.NewGetIntradayHandler(s, log))
		router.Use(server.Pagination)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

}
