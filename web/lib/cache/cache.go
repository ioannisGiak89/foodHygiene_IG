package cache

import (
	"html/template"
	"time"

	"github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"
)

// MemoryCache is caching templates and some of the fsa responses
type MemoryCache struct {
	AuthoritiesResponse *fsaResponse.AuthoritiesResponse
	Templates           *template.Template
}

// NewMemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

// Init the routine that cleans the response data cache
func (mc *MemoryCache) Init() {
	go mc.clearCache()
}

// GetCachedAuthoritiesResponse returns cached response
func (mc *MemoryCache) GetCachedAuthoritiesResponse() *fsaResponse.AuthoritiesResponse {
	return mc.AuthoritiesResponse
}

// SaveAuthoritiesResponse saves the response to memory
func (mc *MemoryCache) SaveAuthoritiesResponse(ar *fsaResponse.AuthoritiesResponse) {
	mc.AuthoritiesResponse = ar
}

// TODO: Add the time that the cache got cleared in order to avoid calling the api 2 times within few seconds
// Clears the cache every 2min
func (mc *MemoryCache) clearCache() {
	for {
		time.Sleep(2 * time.Minute)
		mc.AuthoritiesResponse = nil
	}
}

// LoadTemplates parses all the templates. It is used as cache
func (mc *MemoryCache) LoadTemplates(path string) {
	mc.Templates = template.Must(template.ParseGlob(path))
}

// GetTemplate pulls a template given it's name
func (mc *MemoryCache) GetTemplate(templateName string) *template.Template {
	return mc.Templates.Lookup(templateName)
}
