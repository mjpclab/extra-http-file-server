package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"mjpclab.dev/ghfs/src/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPerms(t *testing.T) {
	mid, err := getPermsUrlsMiddleware([][]string{
		{"/foo", "upload,MKDIR", "user1", "user2"},
	})
	if err != nil {
		t.FailNow()
	}

	users := user.NewList()
	users.AddPlain("user1", "")
	users.AddPlain("user2", "")
	users.AddPlain("user3", "")
	canUpload := false
	canMkdir := false
	canDelete := false
	canArchive := false
	context := &middleware.Context{
		Users:      users,
		CanUpload:  &canUpload,
		CanMkdir:   &canMkdir,
		CanDelete:  &canDelete,
		CanArchive: &canArchive,
	}

	var w *httptest.ResponseRecorder
	var r *http.Request
	var result middleware.ProcessResult

	w = httptest.NewRecorder()

	r = httptest.NewRequest(http.MethodGet, "/abc", nil)
	context.VhostReqPath = r.URL.Path
	result = mid(w, r, context)
	if result != middleware.GoNext {
		t.Error(result)
	}
	if *context.CanUpload {
		t.Error()
	}
	if *context.CanMkdir {
		t.Error()
	}
	if *context.CanDelete {
		t.Error()
	}
	if *context.CanArchive {
		t.Error()
	}

	r = httptest.NewRequest(http.MethodGet, "/foo/bar", nil)
	context.VhostReqPath = r.URL.Path
	result = mid(w, r, context)
	if result != middleware.GoNext {
		t.Error(result)
	}
	if *context.CanUpload {
		t.Error()
	}
	if *context.CanMkdir {
		t.Error()
	}
	if *context.CanDelete {
		t.Error()
	}
	if *context.CanArchive {
		t.Error()
	}

	context.AuthSuccess = true
	context.AuthUserName = "user3"
	result = mid(w, r, context)
	if result != middleware.GoNext {
		t.Error(result)
	}
	if *context.CanUpload {
		t.Error()
	}
	if *context.CanMkdir {
		t.Error()
	}
	if *context.CanDelete {
		t.Error()
	}
	if *context.CanArchive {
		t.Error()
	}

	context.AuthSuccess = true
	context.AuthUserName = "user2"
	result = mid(w, r, context)
	if result != middleware.GoNext {
		t.Error(result)
	}
	if !*context.CanUpload {
		t.Error()
	}
	if !*context.CanMkdir {
		t.Error()
	}
	if *context.CanDelete {
		t.Error()
	}
	if *context.CanArchive {
		t.Error()
	}
	*context.CanUpload = false
	*context.CanMkdir = false

	context.AuthSuccess = true
	context.AuthUserName = "user1"
	result = mid(w, r, context)
	if result != middleware.GoNext {
		t.Error(result)
	}
	if !*context.CanUpload {
		t.Error()
	}
	if !*context.CanMkdir {
		t.Error()
	}
	if *context.CanDelete {
		t.Error()
	}
	if *context.CanArchive {
		t.Error()
	}
	*context.CanUpload = false
	*context.CanMkdir = false
}
