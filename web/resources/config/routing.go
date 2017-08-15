package config

import (
	"net/http"
)

// Route holds data for a route
type Route struct {
	Path    string
	Handler http.HandlerFunc
}

// BuildRoutes takes pre-configured handlers, and assigns them to routes.
func BuildRoutes(services HandlersAsServices) []Route {
	return []Route{
		{"/", services.RedirectHandler},
		{"/home", services.HomeHandler},
		{"/ajax/api/establishments", services.AjaxHomeHandler},
	}
}
