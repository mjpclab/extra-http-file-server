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

		result = middleware.Outputted
		targetUrl, err := url.Parse(target)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		targetUrl = r.URL.ResolveReference(targetUrl)
		if (len(targetUrl.Host) == 0 || targetUrl.Host == r.Host) && targetUrl.RequestURI() == r.RequestURI {
			util.LogErrorString(context.Logger, "redirect to self URL")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, targetUrl.String(), code)
		}
		return
	}, nil
}
