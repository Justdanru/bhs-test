package factory

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/middleware"
	"github.com/google/wire"
)

var middlewaresSet = wire.NewSet(
	middleware.NewInitMiddleware,
	middleware.NewAuthMiddleware,
	middleware.NewRootMiddleware,
)
