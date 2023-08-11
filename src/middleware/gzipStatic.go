package middleware

import (
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/acceptHeaders"
	"mjpclab.dev/ghfs/src/middleware"
	ghfsUtil "mjpclab.dev/ghfs/src/util"
	"net/http"
	"os"
	"strconv"
)

func tryReplaceWithGzFileInfo(w http.ResponseWriter, r *http.Request, context *middleware.Context) (success bool) {
	if len(w.Header().Get("Content-Encoding")) > 0 {
		return
	}

	acceptEncoding := r.Header.Get("Accept-Encoding")
	if len(acceptEncoding) == 0 {
		return
	}
	accepts := acceptHeaders.ParseAccepts(acceptEncoding)
	_, _, canUseGzip := accepts.GetPreferredValue([]string{"gzip"})
	if !canUseGzip {
		return
	}

	gzFsPath := context.AliasFsPath + ".gz"
	file, err := os.Open(gzFsPath)
	if err != nil {
		if os.IsExist(err) {
			util.LogError(context.Logger, err)
		}
		return
	}
	info, err := file.Stat()
	if err != nil {
		util.LogError(context.Logger, err)
		file.Close()
		return
	}
	if info.IsDir() {
		file.Close()
		return
	}
	info = util.CreateRenamedFileInfo((*context.FileInfo).Name(), info)

	(*context.File).Close()
	*context.File = file
	*context.FileInfo = info
	context.AliasFsPath = gzFsPath
	return true
}

func getGzipStaticMiddleware() (middleware.Middleware, error) {
	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		result = middleware.GoNext

		if context.WantJson || !context.AllowAccess || context.File == nil || context.FileInfo == nil || (*context.FileInfo).IsDir() {
			return
		}

		contentType, err := ghfsUtil.GetContentType(context.AliasFsPath, *context.File)
		if err != nil {
			util.LogError(context.Logger, err)
			return
		}

		if !tryReplaceWithGzFileInfo(w, r, context) {
			return
		}

		header := w.Header()
		header.Set("Content-Type", contentType)
		header.Set("Content-Length", strconv.FormatInt((*context.FileInfo).Size(), 10))
		header.Set("Content-Encoding", "gzip")

		return
	}, nil
}
