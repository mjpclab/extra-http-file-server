package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/url"
	"regexp"
)

func rewriteUrl(r *http.Request, targetUrl *url.URL) {
	originalUrl := r.URL
	prefixLen := len(originalUrl.RawPath) - len(originalUrl.Path)
	if prefixLen < 0 {
		prefixLen = 0
	} else if prefixLen > len(originalUrl.RawPath) {
		prefixLen = len(originalUrl.RawPath)
	}
	prefix := originalUrl.RawPath[:prefixLen]

	targetUrl = originalUrl.ResolveReference(targetUrl)
	if len(prefix) > 1 {
		targetUrl.RawPath = prefix + targetUrl.Path
	} else {
		targetUrl.RawPath = targetUrl.Path
	}

	r.URL = targetUrl
}

func getRewriteHostMiddleware(arg [2]string, rewrittenResult middleware.ProcessResult) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var replace string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	replace = arg[1]

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		requestURI := r.Host + r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return middleware.GoNext
		}

		targetUrl, err := util.ReplaceUrl(reMatch, requestURI, replace)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Outputted
		} else {
			rewriteUrl(r, targetUrl)
			return rewrittenResult
		}
	}, nil
}

func getRewriteMiddleware(arg [2]string, rewrittenResult middleware.ProcessResult) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var replace string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	replace = arg[1]

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return middleware.GoNext
		}

		targetUrl, err := util.ReplaceUrl(reMatch, requestURI, replace)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Outputted
		} else {
			rewriteUrl(r, targetUrl)
			return rewrittenResult
		}
	}, nil
}
