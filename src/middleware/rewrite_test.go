package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRewrite(t *testing.T) {
	mid, err := getRewriteMiddleware([2]string{`^/doc(/.*)?`, "/api$1"})
	if err != nil {
		t.FailNow()
	}

	var w *httptest.ResponseRecorder
	var r *http.Request
	var result middleware.ProcessResult

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/abc", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.GoNext {
		t.Error(result)
	}
	if w.Code != 200 {
		t.Error(w.Code)
	}
	if r.URL.Path != "/abc" {
		t.Error(r.URL.Path)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/doc", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.SkipRests {
		t.Error(result)
	}
	if w.Code != 200 {
		t.Error(w.Code)
	}
	if r.URL.Path != "/api" {
		t.Error(r.URL.Path)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/doc/net/http", nil)
	r.URL.RawPath = "/foo/bar/doc/net/http"
	result = mid(w, r, &middleware.Context{})
	if result != middleware.SkipRests {
		t.Error(result)
	}
	if w.Code != 200 {
		t.Error(w.Code)
	}
	if r.URL.Path != "/api/net/http" {
		t.Error(r.URL.Path)
	}
	if r.URL.RawPath != "/foo/bar/api/net/http" {
		t.Error(r.URL.RawPath)
	}

}
