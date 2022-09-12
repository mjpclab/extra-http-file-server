package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnStatus(t *testing.T) {
	mid, err := getReturnStatusMiddleware([2]string{`/doc`, "404"}, nil)
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

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/doc", nil)
	result = mid(w, r, &middleware.Context{})
	if result != middleware.Outputted {
		t.Error(result)
	}
	if w.Code != 404 {
		t.Error(w.Code)
	}
}
