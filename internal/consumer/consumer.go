package consumer

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	reader       *kafka.Reader
	deserializer *protobuf.Deserializer
	processor    messages.IntradayValueProcessor
	logger       *logrus.Entry
}

func NewKafkaConsumer(
	reader *kafka.Reader,
	deserializer *protobuf.Deserializer,
	processor messages.IntradayValueProcessor,
	logger *logrus.Entry,
) *KafkaConsumer {
	return &KafkaConsumer{
		reader:       reader,
		deserializer: deserializer,
		processor:    processor,
		logger:       logger,
	}
}

func (kc *KafkaConsumer) Run(ctx context.Context) {
	for {
		msg, err := kc.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				kc.logger.Info("Context cancelled. Stopping consumer...")
				break
			}
			kc.logger.WithError(err).Error("Failed to read from kafka")
			continue
		}

		err = kc.processMessage(ctx, msg)
		if err != nil {
			kc.logger.WithError(err).Error("Failed to process message")
			continue
		}

		err = kc.reader.CommitMessages(ctx, msg)
		if err != nil {
			kc.logger.WithError(err).Error("Failed to commit message")
		}
	}
}

func (kc *KafkaConsumer) processMessage(ctx context.Context, msg kafka.Message) error {
	kc.logger.Infof("received message: %s", string(msg.Key))
	/**
	   * TODO:
		 * - deserialize protobuf message
		 * - insert data into postgres table
	*/

	value, err := kc.deserializer.Deserialize(msg.Topic, msg.Value)
	if err != nil {
		return fmt.Errorf("unable to deserialize message: %s", err)
	}

	v := value.(*messages.IntradayValue)

	err = kc.processor(ctx, v)

	return nil
}
