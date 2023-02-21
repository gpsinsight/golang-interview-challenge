package messages

import (
	"context"
)

type IntradayValueProcessor func(context.Context, *IntradayValue) error

func NewIntradayValueProcessor(store IntradayValueStore) IntradayValueProcessor {
	return func(ctx context.Context, val *IntradayValue) error {
		return store.Insert(ctx, val)
	}
}
