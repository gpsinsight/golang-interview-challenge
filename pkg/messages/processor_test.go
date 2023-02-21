package messages_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gpsinsight/go-interview-challenge/pkg/messages"
	"github.com/gpsinsight/go-interview-challenge/pkg/messages/messagesfakes"
	"github.com/stretchr/testify/require"
)

func Test_Processor(t *testing.T) {
	ctx := context.TODO()
	store := &messagesfakes.FakeIntradayValueStore{}

	value := &messages.IntradayValue{
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
				store.InsertReturns(nil)
			},
			err: nil,
		},
		{
			label: "store fails",
			setup: func(t *testing.T) {
				store.InsertReturns(fmt.Errorf("unable to write value to store"))
			},
			err: fmt.Errorf("unable to write value to store"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.label, func(t *testing.T) {
			tt.setup(t)
			processor := messages.NewIntradayValueProcessor(store)
			err := processor(ctx, value)
			if err != nil {
				require.EqualError(t, tt.err, err.Error())
			}
		})
	}
}
