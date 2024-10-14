package factory

import (
	"github.com/google/wire"
	"log/slog"
	"os"
)

var loggersSet = wire.NewSet(
	provideLogger,
)

func provideLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
