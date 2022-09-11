package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func getProxyMiddleware(arg [2]string) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var replace string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	replace = arg[1]

	proxy := &httputil.ReverseProxy{
		Director: func(request *http.Request) {},
	}

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
		if err != nil ||
			((len(u.Host) == 0 || u.Host == r.Host) && u.RequestURI() == r.RequestURI) {
			util.LogErrorString(context.Logger, "proxy to self URL")
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Processed
		}

		u = r.URL.ResolveReference(u)
		rProxy, err := http.NewRequest(r.Method, u.String(), r.Body)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return middleware.Processed
		}

		rProxy.Header = r.Header
		proxy.ServeHTTP(w, rProxy)
		return middleware.Processed
	}, nil
}
