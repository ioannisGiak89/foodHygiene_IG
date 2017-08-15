package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/factory"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/service"
	"github.com/ioannisGiak89/foodHygiene_IG/web/model"
)

// AjaxHomeHandler handles ajax requests
type AjaxHomeHandler struct {
	StandardFactory factory.StandardFactory
	FsaService      service.HygieneRatingSystemService
	Cache           *cache.MemoryCache
}

// NewAjaxHomeHandler creates a new AJAX handler
func NewAjaxHomeHandler(sf factory.StandardFactory, fs service.HygieneRatingSystemService, c *cache.MemoryCache) *AjaxHomeHandler {
	return &AjaxHomeHandler{
		StandardFactory: sf,
		FsaService:      fs,
		Cache:           c,
	}
}

// Percentages handles the percentages ajax request
func (ajaxHomeHandler *AjaxHomeHandler) EstablishmentHandler(w http.ResponseWriter, r *http.Request) {
	authorityId := r.FormValue("authority-id")
	if r.Method != "POST" || authorityId == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	establishmentsResponse, err := ajaxHomeHandler.FsaService.GetEstablishments(authorityId)

	if err != nil {
		ajaxHomeHandler.sendResponse(w, "Error", http.StatusInternalServerError)
		// TODO: Log error somewhere useful
		fmt.Println(err)
		return
	}

	calculator := ajaxHomeHandler.StandardFactory.BuildCalculator()
	schemeType := establishmentsResponse.Establishments[0].SchemeType

	var tpl bytes.Buffer
	var calculatedPercentages interface{}
	template := ""

	if schemeType == model.FHRS {
		calculatedPercentages = calculator.CalculateFhrsPercentages(
			establishmentsResponse.Establishments,
		)
		template = "fhrs-table.html"
	} else {
		calculatedPercentages = calculator.CalculateFhisPercentages(
			establishmentsResponse.Establishments,
		)
		template = "fhis-table.html"
	}

	partial := ajaxHomeHandler.Cache.GetTemplate(template)
	err = partial.Execute(&tpl, calculatedPercentages)

	if err != nil {
		ajaxHomeHandler.sendResponse(w, "Error", http.StatusInternalServerError)
		// TODO: Log error somewhere useful
		fmt.Println(err)
		return
	}

	ajaxHomeHandler.sendResponse(w, tpl.String(), http.StatusOK)
}

// sendResponse sends a respond
func (ajaxHomeHandler *AjaxHomeHandler) sendResponse(w http.ResponseWriter, responseBody string, httpStatus int) {
	response := model.AjaxResponse{
		Status: httpStatus,
		Body:   responseBody,
	}

	w.Header().Set("Content-Type", "json")
	w.WriteHeader(httpStatus)

	jsonEncoder := json.NewEncoder(w)
	err := jsonEncoder.Encode(response)
	if err != nil {
		// TODO: Log error somewhere useful
		fmt.Println(err)
	}
	return
}
