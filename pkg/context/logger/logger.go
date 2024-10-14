package logger

import (
	"context"
	"errors"
	"log/slog"
)

var (
	ErrMissingLoggerInContext = errors.New("missing logger in context")
)

type loggerKey struct{}

func ContextWithLogger(parent context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(parent, loggerKey{}, logger)
}

func FromContext(ctx context.Context) (*slog.Logger, error) {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger, nil
	}

	return nil, ErrMissingLoggerInContext
}
