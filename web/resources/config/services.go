package config

import (
	"net/http"
)

// Services contains all the handlers as services.
type HandlersAsServices struct {
	// Handlers
	HomeHandler     http.HandlerFunc
	AjaxHomeHandler http.HandlerFunc
	RedirectHandler http.HandlerFunc
}
