package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ioannisGiak89/foodHygiene_IG/web/handler"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/factory"
	"github.com/ioannisGiak89/foodHygiene_IG/web/resources/config"
)

func main() {

	// Init the app
	standardFactory := factory.NewAppFactory()
	fsaClient := standardFactory.BuildFsaClient()
	cache := standardFactory.BuildCache()
	// Starts the routine for clearing the cache
	cache.Init()
	fsaService := standardFactory.BuildFsaService(fsaClient, cache)

	// Caching the templates
	cache.LoadTemplates("web/resources/views/*/*")

	// Init the file server for static files
	handler.NewFileServerHandler().Init()

	// Load handler services
	handlers := config.HandlersAsServices{
		HomeHandler:     handler.NewHomeHandler(cache, fsaService).HomeHandler,
		AjaxHomeHandler: handler.NewAjaxHomeHandler(standardFactory, fsaService, cache).EstablishmentHandler,
		RedirectHandler: handler.NewRedirectHandler().RedirectHandler,
	}

	// Routing
	for _, route := range config.BuildRoutes(handlers) {
		http.HandleFunc(route.Path, route.Handler)
	}

	fmt.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
