package messages

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . IntradayValueStore
type IntradayValueStore interface {
	Insert(ctx context.Context, value *IntradayValue) error
	List(ctx context.Context, limit int, offset int) ([]*IntradayValue, error)
}
