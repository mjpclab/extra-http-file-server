package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"regexp"
	"strconv"
)

func getReturnStatusMiddleware(arg [2]string, statusPageMids []middleware.Middleware) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var code int

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	code, err = strconv.Atoi(arg[1])
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return middleware.SkippedGoNext
		}

		result = middleware.Outputted
		context.Status = code
		for i := range statusPageMids {
			if statusPageMids[i](w, r, context) == middleware.Outputted {
				return
			}
		}

		w.WriteHeader(code)
		return
	}, nil
}
