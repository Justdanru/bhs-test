package factory

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/handler"
	"github.com/google/wire"
)

var handlersSet = wire.NewSet(
	handler.NewErrorsHandler,
	handler.NewUserHandler,
	handler.NewCheckUsernameHandler,
	handler.NewRegisterHandler,
	handler.NewRootHandler,
)
