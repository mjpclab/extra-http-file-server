package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"regexp"
)

func getHeaderAddMiddleware(arg [3]string) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var name, value string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	name = arg[1]
	value = arg[2]

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		result = middleware.GoNext

		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return
		}

		w.Header().Add(name, value)
		return
	}, nil
}

func getHeaderSetMiddleware(arg [3]string) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var name, value string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	name = arg[1]
	value = arg[2]

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		result = middleware.GoNext

		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return
		}

		w.Header().Set(name, value)
		return
	}, nil
}
