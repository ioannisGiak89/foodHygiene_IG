package handler

import (
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/service"
	"github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"
)

type FakeFsaService struct{}

func NewFakeFsaService() service.HygieneRatingSystemService {
	return &FakeFsaService{}
}

func (fs *FakeFsaService) GetAuthorities() (*fsaResponse.AuthoritiesResponse, error) {
	return authoritiesResponse, nil
}

func (fs *FakeFsaService) GetEstablishments(authorityId string) (*fsaResponse.EstablishmentsResponse, error) {
	return establishmentResponse, nil
}
