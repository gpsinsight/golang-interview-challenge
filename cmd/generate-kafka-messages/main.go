package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"github.com/gpsinsight/go-interview-challenge/internal/config"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/kelseyhightower/envconfig"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

//go:embed testdata/stock-values.csv
var stockData []byte

func main() {
	var cfg config.Config
	_ = envconfig.Process("", &cfg)

	log := getLogger()

	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka.Brokers...),
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.Hash{},
		Topic:                  cfg.Kafka.Topics.IntradayValues,
		Async:                  true,
	}
	defer w.Close()

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(cfg.Kafka.SchemaRegistry.URL))
	if err != nil {
		log.WithError(err).Fatal("not connected to schema registry")
	}
	serConfig := protobuf.NewSerializerConfig()
	serConfig.AutoRegisterSchemas = true
	protoSerializer, err := protobuf.NewSerializer(client, serde.ValueSerde, serConfig)
	if err != nil {
		log.WithError(err).Error("failed to set up serializer")
		return
	}

	data, err := parseCSV(stockData)
	if err != nil {
		log.Fatalf("Failed to parse csv data: %s", err)
	}

	for _, value := range data {
		payload, err := protoSerializer.Serialize(cfg.Kafka.Topics.IntradayValues, value)
		if err != nil {
			log.WithError(err).Fatal("failed to serialize")
		}

		msg := kafka.Message{
			Key:   []byte(value.Ticker),
			Value: payload,
		}

		w.WriteMessages(context.Background(), msg)
	}
}

func parseCSV(data []byte) ([]*messages.IntradayValue, error) {
	// Create a CSV reader
	reader := csv.NewReader(bytes.NewReader(data))

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Check if there are at least a header and one row
	if len(rows) < 2 {
		return nil, fmt.Errorf("no data rows found in CSV")
	}

	// Skip the header row
	rows = rows[1:]

	var stockDataList []*messages.IntradayValue
	for _, row := range rows {
		if len(row) != 7 {
			return nil, fmt.Errorf("invalid row length: %v", row)
		}

		// Parse the fields
		ticker := row[0]

		timestamp, err := time.Parse("2006-01-02 15:04:05", row[1])
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp format: %v", row[1])
		}

		open, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid open value: %v", row[2])
		}

		high, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid high value: %v", row[3])
		}

		low, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid low value: %v", row[4])
		}

		closePrice, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid close value: %v", row[5])
		}

		volume, err := strconv.Atoi(row[6])
		if err != nil {
			return nil, fmt.Errorf("invalid volume value: %v", row[6])
		}

		// Append the parsed data to the list
		stockDataList = append(stockDataList, &messages.IntradayValue{
			Ticker:    ticker,
			Timestamp: timestamp.Format(time.RFC3339Nano),
			Open:      open,
			High:      high,
			Low:       low,
			Close:     closePrice,
			Volume:    int64(volume),
		})
	}

	return stockDataList, nil
}

func getLogger() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log := logger.WithField("environment", "local")

	return log
}
