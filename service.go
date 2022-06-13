package lunobot

import (
	"context"
)

type Service func(ctx context.Context, s *bot, m *Message) error
