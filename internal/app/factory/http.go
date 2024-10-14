package factory

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/router"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/server"
	"github.com/google/wire"
	"net/http"
)

var httpSet = wire.NewSet(
	router.NewRouter,
	wire.Bind(new(http.Handler), new(*router.Router)),

	server.NewHTTPServer,
)
