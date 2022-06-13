package lunobot

import (
	"context"
)

const (
	timePeriodOneDay = "1D"

	optionMarketCap = "mktcap"
)

func statsHandler(ctx context.Context, b *bot, m *Message) error {
	switch m.Trailing[0] {
	case optionMarketCap:
		return getMarketCap(ctx, b, m.Trailing[1])
	default:
		panic("not implemented")
	}
}

func getMarketCap(ctx context.Context, b *bot, token string) error {
	return getMarketCapTimePeriod(ctx, b, token, timePeriodOneDay, false)
}

func getMarketCapTimePeriod(ctx context.Context, b *bot, token string, period string, chart bool) error {
	return nil
}
