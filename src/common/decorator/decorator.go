package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

type CommandDecorator[C any] interface {
	Handle(ctx context.Context, c C) error
}

func generateActionName(handler any) string {

	return strings.Split(fmt.Sprintf("%T", handler), ".")[0]
}

func ApplyCommandDecorator[C any](
	handler CommandDecorator[C],
	logger *slog.Logger,
) CommandDecorator[C] {
	return CommandLoggingDecorator[C]{
		logger: logger,
		base:   handler,
	}
}
