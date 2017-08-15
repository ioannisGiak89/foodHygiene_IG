package handler

import (
	"net/http"

	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/service"
)

// HomeHandler handles home page
type HomeHandler struct {
	FsaService service.HygieneRatingSystemService
	Cache      *cache.MemoryCache
}

// Creates a new Home Handler
func NewHomeHandler(c *cache.MemoryCache, fs service.HygieneRatingSystemService) *HomeHandler {
	return &HomeHandler{
		Cache:      c,
		FsaService: fs,
	}
}

// HomeHandler handles http calls to /home
func (h *HomeHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	authorities, err := h.FsaService.GetAuthorities()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Cache.Templates.ExecuteTemplate(w, "home.html", authorities.Authorities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
