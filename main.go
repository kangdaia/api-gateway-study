package main

import (
	"api-gateway-study/app"
	"api-gateway-study/app/dependency"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		dependency.Cfg,
		dependency.HttpClient,
		dependency.Producer,
		dependency.Router,
		fx.Provide(app.NewApp),
		fx.Invoke(func(app.App) {}),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),
	).Run()
}
