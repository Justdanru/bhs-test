// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package factory

import (
	"github.com/Justdanru/bhs-test/internal/app"
)

// Injectors from wire.go:

func StartApp() (*app.App, func(), error) {
	appApp, cleanup, err := startApp()
	if err != nil {
		return nil, nil, err
	}
	return appApp, func() {
		cleanup()
	}, nil
}
