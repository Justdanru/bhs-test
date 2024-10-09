package app

import "fmt"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	fmt.Printf("App started\n")
	return nil
}

func (a *App) Shutdown() {
	fmt.Printf("App closed\n")
}
