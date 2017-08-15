package service

import (
	"encoding/json"

	"github.com/ioannisGiak89/foodHygiene_IG/web/lib"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
	"github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"
)

// HygieneRatingSystemService defines the Hygiene Rating System Service
type HygieneRatingSystemService interface {
	// GetAuthorities gets a list with all the authorities and some basic information
	GetAuthorities() (*fsaResponse.AuthoritiesResponse, error)
	// GetEstablishments gets all the Establishments for a given authority id
	GetEstablishments(authorityId string) (*fsaResponse.EstablishmentsResponse, error)
}

// FsaService implements the HygieneRatingSystemService interface
type FsaService struct {
	HttpClient *lib.HttpClient
	Cache      *cache.MemoryCache
}

// NewFsaService creates a new FsaService
func NewFsaService(httpClient *lib.HttpClient, cache *cache.MemoryCache) HygieneRatingSystemService {
	return &FsaService{
		HttpClient: httpClient,
		Cache:      cache,
	}
}

func (fs *FsaService) GetAuthorities() (*fsaResponse.AuthoritiesResponse, error) {
	cachedResponse := fs.Cache.GetCachedAuthoritiesResponse()
	if cachedResponse != nil {
		return cachedResponse, nil
	}

	responseBody, err := fs.HttpClient.Get("Authorities/basic")

	if err != nil {
		return nil, err
	}

	var authorities fsaResponse.AuthoritiesResponse
	err = json.Unmarshal(responseBody, &authorities)

	if err != nil {
		return nil, err
	}

	// Cache Authorities and return response
	fs.Cache.SaveAuthoritiesResponse(&authorities)
	return &authorities, nil
}

func (fs *FsaService) GetEstablishments(authorityId string) (*fsaResponse.EstablishmentsResponse, error) {

	responseBody, err := fs.HttpClient.Get("Establishments?pageSize=0&localAuthorityId=" + authorityId)

	if err != nil {
		return nil, err
	}

	var establishments fsaResponse.EstablishmentsResponse
	err = json.Unmarshal(responseBody, &establishments)

	if err != nil {
		return nil, err
	}

	return &establishments, nil
}
