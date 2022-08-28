package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func getRewriteMiddleware(arg [2]string) (middleware.Middleware, error) {
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

		u, err := url.Parse(target)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Processed
		} else {
			originalUrl := r.URL
			prefixLen := len(originalUrl.RawPath) - len(originalUrl.Path)
			if prefixLen < 0 {
				prefixLen = 0
			} else if prefixLen > len(originalUrl.RawPath) {
				prefixLen = len(originalUrl.RawPath)
			}
			prefix := originalUrl.RawPath[:prefixLen]

			r.URL = originalUrl.ResolveReference(u)
			if len(prefix) > 1 { // if prefix=="/", skip
				r.URL.RawPath = prefix + r.URL.Path
			} else {
				r.URL.RawPath = r.URL.Path
			}

			return middleware.SkipRests
		}
	}, nil
}
