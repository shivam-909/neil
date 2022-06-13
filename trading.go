package lunobot

import (
	"context"
	"strconv"
)

const (
	sideBuy  = "buy"
	sideSell = "sell"
)

func tradingHandler(ctx context.Context, b *bot, m *Message) error {
	switch m.Trailing[0] {
	case sideBuy:
		amount, err := strconv.Atoi(m.Trailing[2])
		if err != nil {
			return err
		}
		return createOrder(ctx, b, sideBuy, m.Trailing[1], amount)
	case sideSell:
		amount, err := strconv.Atoi(m.Trailing[2])
		if err != nil {
			return err
		}
		return createOrder(ctx, b, sideSell, m.Trailing[1], amount)
	default:
		return nil
	}
}

func createOrder(ctx context.Context, b *bot, side string, token string, amount int) error {
	return nil
}
