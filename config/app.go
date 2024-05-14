package config

import (
	"os"
	"starter_kit/public"

	"github.com/caesar-rocks/core"
)

func ProvideApp(env *EnvironmentVariables) *core.App {
	app := core.NewApp(&core.AppConfig{
		Addr: os.Getenv("ADDR"),
	})

	// Register controllers
	app.RegisterProviders(
	// ...
	)

	// Register the repositories
	app.RegisterProviders(
	// ...
	)

	// Register the rest
	app.RegisterProviders(
		RegisterRoutes,
		ProvideEnvironmentVariables,
		ProvideDatabase,
		ProvideErrorHandler,
		// ...
	)

	app.RegisterInvokers(
		core.ServeStaticAssets(public.FS),
		// ...
	)

	return app
}
