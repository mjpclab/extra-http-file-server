package middleware

import (
	"crypto/tls"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
)

func isPkiValidationResource(requestPath string) bool {
	const valiPath = "/.well-known/pki-validation/"
	return len(requestPath) > len(valiPath) && requestPath[:len(valiPath)] == valiPath
}

func getPkiValidationSkipToHttpsMiddleware() middleware.Middleware {
	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		result = middleware.GoNext

		if r.TLS != nil {
			return
		}

		// skip https redirect for special url /.well-known/pki-validation/
		// set `Request.TLS` a value to cheat redirect logic skipping redirect
		if isPkiValidationResource(r.URL.Path) {
			connState := tls.ConnectionState{}
			r.TLS = &connState
		}
		return
	}
}
