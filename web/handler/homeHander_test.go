package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/ioannisGiak89/foodHygiene_IG/web/lib/cache"
)

func Test_HomeHandler(t *testing.T) {
	t.Run("should return home.html template", func(t *testing.T) {
		memCache := cache.NewMemoryCache()
		memCache.LoadTemplates("../resources/views/*/*")
		homeHandler := NewHomeHandler(memCache, NewFakeFsaService()).HomeHandler
		req, err := http.NewRequest("GET", "/home", nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		homeHandler(res, req)
		exp := "Aberdeenshire"
		act := res.Body.String()
		if !strings.Contains(act, exp) {
			t.Fatalf("Template doesn't contain %s", exp)
		}
	})
}
