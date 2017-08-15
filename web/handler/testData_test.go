package handler

import "github.com/ioannisGiak89/foodHygiene_IG/web/model/fsaResponse"

var authoritiesResponse = &fsaResponse.AuthoritiesResponse{
	Authorities: []fsaResponse.Authority{
		{LocalAuthorityId: 197, Name: "Aberdeen City"},
		{LocalAuthorityId: 198, Name: "Aberdeenshire"},
		{LocalAuthorityId: 277, Name: "Adur"},
	},
}

var establishmentResponse = &fsaResponse.EstablishmentsResponse{
	Establishments: []fsaResponse.Establishment{
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Pass and Eat Safe", SchemeType: "FHIS"},
		{RatingValue: "Pass and Eat Safe", SchemeType: "FHIS"},
		{RatingValue: "Pass and Eat Safe", SchemeType: "FHIS"},
		{RatingValue: "Improvement Required", SchemeType: "FHIS"},
		{RatingValue: "Awaiting Inspection", SchemeType: "FHIS"},
		{RatingValue: "Exempt", SchemeType: "FHIS"},
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Awaiting Inspection", SchemeType: "FHIS"},
	},
}
