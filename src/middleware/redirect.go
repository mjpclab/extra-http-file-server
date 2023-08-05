package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"regexp"
	"strconv"
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

		result = middleware.Outputted
		targetUrl, err := util.ReplaceUrl(reMatch, requestURI, replace)
		if err != nil {
			util.LogError(context.Logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if prefixLen := len(context.PrefixReqPath) - len(context.VhostReqPath); prefixLen > 0 && len(targetUrl.Host) == 0 {
			prefix := context.PrefixReqPath[:prefixLen]
			targetUrl.Path = prefix + targetUrl.Path
		}
		targetUrl = r.URL.ResolveReference(targetUrl)
		if util.IsUrlSameAsReq(targetUrl, r) {
			util.LogErrorString(context.Logger, "redirect to self URL")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, targetUrl.String(), code)
		}
		return
	}, nil
}
