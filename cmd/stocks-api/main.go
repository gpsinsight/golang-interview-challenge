package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"github.com/gpsinsight/go-interview-challenge/internal/config"
	"github.com/gpsinsight/go-interview-challenge/internal/consumer"
	"github.com/gpsinsight/go-interview-challenge/internal/server"
	"github.com/gpsinsight/go-interview-challenge/internal/stocks"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := getLogger()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("config.New", err)
	}

	ctx := getContext(logger)

	db := sqlx.MustConnect("postgres", cfg.Postgres.ConnectionString())
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(time.Hour)
	defer db.Close()

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.Kafka.Brokers,
		GroupID:     cfg.Kafka.GroupID,
		GroupTopics: []string{cfg.Kafka.Topics.IntradayValues},
		ErrorLogger: kafka.LoggerFunc(logger.Errorf),
	})
	defer kafkaReader.Close()

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(cfg.Kafka.SchemaRegistry.URL))
	if err != nil {
		logger.WithError(err).Fatal("not connected to schema registry")
	}

	protoDeserializer, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		logger.WithError(err).Fatal("failed to set up deserializer")
	}
	err = protoDeserializer.ProtoRegistry.RegisterMessage((&messages.IntradayValue{}).ProtoReflect().Type())
	if err != nil {
		logger.WithError(err).Fatal("failed to register IntradayValue message type")
	}

	stocksRepository := stocks.NewRepository(logger, db)

	stocksService := stocks.NewService(logger, stocksRepository)

	kafkaConsumer := consumer.NewKafkaConsumer(
		kafkaReader,
		protoDeserializer,
		stocksService,
		logger,
	)
	go kafkaConsumer.Run(ctx)

	logger.Info("Starting up go-interview-challenge")

	srvr := server.New(cfg, logger, stocksService)
	defer func() {
		err := srvr.Shutdown(context.Background())
		if err != nil {
			logger.WithError(err).Error("failed to gracefully shutdown server")
		}
	}()
	go func() {
		err := srvr.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err)
		}
	}()

	// Exit safely
	<-ctx.Done()
	logger.Info("exiting")
}

func getLogger() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	log := logger.WithField("environment", "local")

	return log
}

func getContext(logger *logrus.Entry) context.Context {
	// Setup context that will cancel on signalled termination
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sig
		logger.Info("termination signaled")
		cancel()
	}()

	return ctx
}
