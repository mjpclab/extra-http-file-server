package middleware

import (
	"crypto/tls"
	"mjpclab.dev/ghfs/src/middleware"
	"mjpclab.dev/ghfs/src/util"
	"net/http"
)

func getPkiValidationSkipToHttpsMiddleware() middleware.Middleware {
	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		result = middleware.GoNext

		if r.TLS != nil {
			return
		}

		// skip https redirect for special url /.well-known/
		// set `Request.TLS` a value to cheat redirect logic skipping redirect
		if util.HasUrlPrefixDir(context.PrefixReqPath, "/.well-known") {
			connState := tls.ConnectionState{}
			r.TLS = &connState
		}
		return
	}
}
