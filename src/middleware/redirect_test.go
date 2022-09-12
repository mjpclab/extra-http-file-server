package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	mid, err := getRedirectMiddleware([3]string{`\?goto=(.*)`, "$1", "307"})
	if err != nil {
		t.FailNow()
	}

	var w *httptest.ResponseRecorder
	var r *http.Request
	var result middleware.ProcessResult
	var location string

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/abc", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.GoNext {
		t.Error(result)
	}
	if w.Code != 200 {
		t.Error(w.Code)
	}
	location = w.Header().Get("Location")
	if location != "" {
		t.Error(location)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/abc?goto=/", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.Outputted {
		t.Error(result)
	}
	if w.Code != 307 {
		t.Error(w.Code)
	}
	location = w.Header().Get("Location")
	if location != "/" {
		t.Error(location)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/abc?goto=http://www.example.com/", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.Outputted {
		t.Error(result)
	}
	if w.Code != 307 {
		t.Error(w.Code)
	}
	location = w.Header().Get("Location")
	if location != "http://www.example.com/" {
		t.Error(location)
	}
}
