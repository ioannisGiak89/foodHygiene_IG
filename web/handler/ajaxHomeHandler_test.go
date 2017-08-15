package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/factory"
)

func TestAjaxHomeHandler_EstablishmentHandler(t *testing.T) {
	t.Run("should return fhis table", func(t *testing.T) {
		memCache := cache.NewMemoryCache()
		memCache.LoadTemplates("../resources/views/*/*")
		AjaxHomeHandler := NewAjaxHomeHandler(factory.NewAppFactory(), NewFakeFsaService(), memCache).EstablishmentHandler
		req, err := http.NewRequest("POST", "/ajax/api/establishments?authority-id=1", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		AjaxHomeHandler(res, req)
		exp := "Pass"
		act := res.Body.String()
		if !strings.Contains(act, exp) {
			t.Fatalf("Template doesn't contain %s", exp)
		}
	})
}
