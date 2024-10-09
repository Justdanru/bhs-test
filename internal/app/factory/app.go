package factory

import (
	"fmt"
	"github.com/Justdanru/bhs-test/internal/app"
)

func startApp() (*app.App, func(), error) {
	newApp := app.NewApp()

	err := newApp.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't run app. %w", err)
	}

	return newApp, func() {
		newApp.Shutdown()
	}, nil
}
