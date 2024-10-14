package factory

import (
	"fmt"
	"github.com/Justdanru/bhs-test/internal/app"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/server"
)

func startApp(
	httpServer *server.HTTPServer,
) (*app.App, func(), error) {
	newApp := app.NewApp(httpServer)

	err := newApp.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't run app. %w", err)
	}

	return newApp, func() {
		newApp.Shutdown()
	}, nil
}
