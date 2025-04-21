package stocks

import (
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Entry
	repo   *Repository
}

func NewService(logger *logrus.Entry, repo *Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func ProcessStockMessage(msg *messages.IntradayValue) error {
	// TODO: insert new stock data into using InsertIntradayValue on repository
	return nil
}
