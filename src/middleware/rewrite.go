package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

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
		matches := reMatch.FindStringSubmatch(requestURI)
		if len(matches) > 10 {
			matches = matches[:10]
		}

		target := replace
		for i := range matches {
			target = strings.ReplaceAll(target, "$"+strconv.Itoa(i), matches[i])
		}
		if len(target) == 0 {
			target = "/"
		}

		targetUrl, err := url.Parse(target)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Outputted
		} else {
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
			return rewrittenResult
		}
	}, nil
}
