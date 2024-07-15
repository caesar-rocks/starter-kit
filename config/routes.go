package config

import (
	"starter_kit/views/pages"

	caesar "github.com/caesar-rocks/core"
)

func authMiddleware(ctx *caesar.Context) error {
	ctx.Next()
	return nil
}

func RegisterRoutes(
// Inject your controllers here...
) *caesar.Router {
	router := caesar.NewRouter()
	router.Render("/", pages.WelcomePage()).Use(authMiddleware)

	return router
}
