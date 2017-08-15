package handler

import (
	"net/http"
)

// RedirectHandler redirects to home page
type RedirectHandler struct{}

// Creates a new Redirect Handler
func NewRedirectHandler() *RedirectHandler {
	return &RedirectHandler{}
}

// RedirectHandler redirects to home page
func (h *RedirectHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}
