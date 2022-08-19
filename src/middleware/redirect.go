package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func getRedirectMiddleware(arg [3]string) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var replace string
	var code int

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	replace = arg[1]
	code, err = strconv.Atoi(arg[2])
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (processed bool) {
		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return false
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
		if err != nil ||
			((len(u.Host) == 0 || u.Host == r.Host) && u.RequestURI() == r.RequestURI) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, u.String(), code)
		}
		return true
	}, nil
}
