package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type CommandLoggingDecorator[C any] struct {
	base   CommandDecorator[C]
	logger *slog.Logger
}

func (d CommandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) error {
	start := time.Now()
	handlerType := generateActionName(cmd)

	d.logger.Debug("Executing command",
		slog.String("command", handlerType),
		slog.String("command_body", fmt.Sprintf("%#v", cmd)),
	)
	err := d.base.Handle(ctx, cmd)
	defer func() {
		duration := time.Since(start)

		if err != nil {
			d.logger.Error("Failed to execute command",
				slog.String("command", handlerType),
				slog.Duration("duration", duration),
				slog.String("error", err.Error()),
			)

		} else {
			d.logger.Info("Command executed successfully",
				slog.String("command", handlerType),
				slog.Duration("duration", duration),
			)
		}
	}()

	return err
}
