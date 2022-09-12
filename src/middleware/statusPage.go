package middleware

import (
	"io"
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"mjpclab.dev/ghfs/src/serverHandler"
	ghfsUtil "mjpclab.dev/ghfs/src/util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func getStatusPageMiddleware(arg [2]string) (middleware.Middleware, error) {
	code, err := strconv.Atoi(arg[0])
	if err != nil {
		return nil, err
	}

	if len(arg[1]) == 0 {
		return nil, errInvalidParamValue
	}

	statusFile := filepath.Clean(arg[1])
	if len(statusFile) == 0 {
		return nil, errInvalidParamValue
	}
	if statusFile[len(statusFile)-1] == '.' { // "." or "c:."
		return nil, errInvalidParamValue
	}

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		if context.Status != code {
			return middleware.GoNext
		}

		file, err := os.Open(statusFile)
		if err != nil {
			util.LogError(context.Logger, err)
			return middleware.GoNext
		}

		contentType, err := ghfsUtil.GetContentType(statusFile, file)
		if err != nil {
			util.LogError(context.Logger, err)
			return middleware.GoNext
		}
		w.Header().Set("Content-Type", contentType)

		w.WriteHeader(context.Status)
		if serverHandler.NeedResponseBody(r.Method) {
			_, err = io.Copy(w, file)
			if err != nil {
				util.LogError(context.Logger, err)
			}
			err = file.Close()
			if err != nil {
				util.LogError(context.Logger, err)
			}
		}

		return middleware.Outputted
	}, nil
}
