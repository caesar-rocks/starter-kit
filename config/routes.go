package config

import (
	"starter_kit/views/pages"

	"github.com/caesar-rocks/core"
)

func RegisterRoutes(
// Inject your controllers here...
) *core.Router {
	router := core.NewRouter()
	router.Render("/", pages.WelcomePage())

	return router
}
