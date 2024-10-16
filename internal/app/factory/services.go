package factory

import (
	"github.com/Justdanru/bhs-test/internal/infrastructure/service/auth"
	"github.com/Justdanru/bhs-test/internal/infrastructure/service/user"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	"github.com/google/wire"
)

var servicesSet = wire.NewSet(
	user.NewService,
	wire.Bind(new(service.UserService), new(*user.Service)),

	auth.NewService,
	wire.Bind(new(service.AuthService), new(*auth.Service)),
)
