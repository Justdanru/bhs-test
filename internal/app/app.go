package app

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/server"
)

type App struct {
	HTTPServer *server.HTTPServer
}

func NewApp(httpServer *server.HTTPServer) *App {
	return &App{
		HTTPServer: httpServer,
	}
}

func (a *App) Run() error {
	return a.HTTPServer.Run()
}

func (a *App) Shutdown() {
	a.HTTPServer.Shutdown()
}
