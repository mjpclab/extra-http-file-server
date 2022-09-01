package middleware

import (
	"io"
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func getReturnStatusMiddleware(arg [3]string) (middleware.Middleware, error) {
	var err error
	var reMatch *regexp.Regexp
	var code int
	var filename string

	reMatch, err = regexp.Compile(arg[0])
	if err != nil {
		return nil, err
	}
	code, err = strconv.Atoi(arg[1])
	if err != nil {
		return nil, err
	}
	filename = arg[2]

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		requestURI := r.URL.RequestURI() // request uri without prefix path
		if !reMatch.MatchString(requestURI) {
			return middleware.GoNext
		}

		result = middleware.Processed

		w.WriteHeader(code)

		if len(filename) > 0 {
			file, err := os.Open(filename)
			if err != nil {
				util.LogError(context.Logger, err)
				return
			}

			_, err = io.Copy(w, file)
			if err != nil {
				util.LogError(context.Logger, err)
			}
			err = file.Close()
			if err != nil {
				util.LogError(context.Logger, err)
			}
		}
		return
	}, nil
}
