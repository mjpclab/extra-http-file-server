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
	var replacement string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	replacement = arg[1]

	proxy := &httputil.ReverseProxy{
		Rewrite: func(proxyReq *httputil.ProxyRequest) {},
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

		targetUrl, err := util.ReplaceUrl(reMatch, requestURI, replacement)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		targetUrl = r.URL.ResolveReference(targetUrl)
		if len(targetUrl.Host) == 0 {
			targetUrl.Host = r.Host
		}
		if util.IsUrlSameAsReq(targetUrl, r) {
			util.LogErrorString(context.Logger, "proxy to self URL")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(targetUrl.Scheme) == 0 {
			if len(r.URL.Scheme) > 0 {
				targetUrl.Scheme = r.URL.Scheme
			} else if r.TLS != nil {
				targetUrl.Scheme = "https"
			} else {
				targetUrl.Scheme = "http"
			}
		}

		outHeader := r.Header.Clone()
		outHeader.Set("Host", targetUrl.Host)
		outHeader.Set("Referer", targetUrl.RequestURI())
		outHeader.Set("Origin", targetUrl.Scheme+"://"+targetUrl.Host)
		outReq := &http.Request{
			Method: r.Method,
			URL:    targetUrl,
			Body:   r.Body,
			Header: outHeader,
		}
		proxy.ServeHTTP(w, outReq)
		go util.LogAccess(context.Logger, "proxy request to "+outReq.URL.String())
		return
	}, nil
}
