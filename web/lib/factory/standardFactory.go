package factory

import (
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/percentages"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/service"
)

// Factory abstracts the creation of services.
type StandardFactory interface {
	// BuildEntryFacade builds an EntryFacade instance.
	BuildCalculator() *percentages.Calculator
	// BuildFsaService builds Fsa Service
	BuildFsaService(*lib.HttpClient, *cache.MemoryCache) service.HygieneRatingSystemService
	// BuildFsaClient builds the Fsa Client
	BuildFsaClient() *lib.HttpClient
	// BuildCache builds the cache
	BuildCache() *cache.MemoryCache
}

// AppFactory builds services
type AppFactory struct{}

// StandardFactory creates a Calculator
func NewAppFactory() StandardFactory {
	return &AppFactory{}
}

func (f *AppFactory) BuildCalculator() *percentages.Calculator {
	return percentages.NewCalculator()
}

func (f *AppFactory) BuildFsaService(client *lib.HttpClient, mc *cache.MemoryCache) service.HygieneRatingSystemService {
	return service.NewFsaService(client, mc)
}

func (f *AppFactory) BuildFsaClient() *lib.HttpClient {
	return lib.NewHttpClient("http://api.ratings.food.gov.uk/")
}

func (f *AppFactory) BuildCache() *cache.MemoryCache {
	return cache.NewMemoryCache()
}
