package messages

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . IntradayStore
type IntradayStore interface {
	Insert(context.Context, IntradayValue) error
	List(context.Context) ([]IntradayValue, error)
}
