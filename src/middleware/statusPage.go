package middleware

import (
	"io"
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"mjpclab.dev/ghfs/src/serverHandler"
	"net/http"
	"path/filepath"
	"strconv"
)

func getStatusPageMiddleware(arg [2]string, enableGzipStatic bool) (middleware.Middleware, error) {
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
		if *context.Status != code {
			return middleware.GoNext
		}

		file, info, contentType, err := util.GetFileInfoType(statusFile)
		if err != nil {
			util.LogError(context.Logger, err)
			return middleware.GoNext
		}
		if context.File != nil {
			(*context.File).Close()
			*context.File = file
		} else {
			context.File = &file
			defer file.Close()
		}
		if context.FileInfo != nil {
			*context.FileInfo = info
		} else {
			context.FileInfo = &info
		}
		context.AliasFsPath = statusFile

		useGzipStatic := enableGzipStatic && tryReplaceWithGzFileInfo(w, r, context)

		header := w.Header()
		header.Set("Last-Modified", (*context.FileInfo).ModTime().UTC().Format(http.TimeFormat))
		header.Set("Content-Type", contentType)
		header.Set("Content-Length", strconv.FormatInt((*context.FileInfo).Size(), 10))
		if useGzipStatic {
			header.Set("Content-Encoding", "gzip")
		}

		w.WriteHeader(code)
		if serverHandler.NeedResponseBody(r.Method) {
			_, err = io.Copy(w, file)
			if err != nil {
				util.LogError(context.Logger, err)
			}
		}

		return middleware.Outputted
	}, nil
}
