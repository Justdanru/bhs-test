package main

import (
	"fmt"
	"github.com/Justdanru/bhs-test/internal/app/factory"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{
		Name: "app",
		Action: func(c *cli.Context) error {
			_, cleanup, err := factory.StartApp()
			if err != nil {
				return fmt.Errorf("error start app. %w", err)
			}

			defer cleanup()

			sig := make(chan os.Signal, 2)
			signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
			<-sig

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
