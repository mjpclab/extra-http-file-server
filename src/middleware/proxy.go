package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"net/http/httputil"
	"regexp"
)

func getProxyMiddleware(arg [2]string, preOutputMids []middleware.Middleware) (middleware.Middleware, error) {
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

		result = middleware.Outputted
		for i := range preOutputMids {
			preOutputMids[i](w, r, context)
		}

		targetUrl, err := util.ReplaceUrl(reMatch, requestURI, replace)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		targetUrl = r.URL.ResolveReference(targetUrl)
		if (len(targetUrl.Host) == 0 || targetUrl.Host == r.Host) && targetUrl.RequestURI() == r.RequestURI {
			util.LogErrorString(context.Logger, "proxy to self URL")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rProxy, err := http.NewRequest(r.Method, targetUrl.String(), r.Body)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rProxy.Header = r.Header
		proxy.ServeHTTP(w, rProxy)
		return
	}, nil
}
