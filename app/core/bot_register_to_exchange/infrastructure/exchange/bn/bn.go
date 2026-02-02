package exchange

import "context"

type IBNExchange interface {
	RegisterBot(ctx context.Context) error
}
