package messages

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . IntradayValueStore
type IntradayValueStore interface {
	Insert(context.Context, *IntradayValue) error
	List(context.Context) ([]IntradayValue, error)
}
