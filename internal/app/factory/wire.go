//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/Justdanru/bhs-test/internal/app"
	"github.com/google/wire"
)

func StartApp() (*app.App, func(), error) {
	panic(wire.Build(
		configSet,
		loggersSet,
		repositoriesSet,
		servicesSet,
		handlersSet,
		httpSet,
		startApp,
	))
}
