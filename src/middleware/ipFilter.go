package middleware

import (
	"bytes"
	"mjpclab.dev/ehfs/src/lib"
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	"net/http"
	"os"
)

func getIPRangeMan(ips, ipFiles []string) (man *lib.IPRangeMan, errs []error) {
	var err error
	man = lib.NewIPRangeMan()

	for i := range ips {
		err = man.AddByString(ips[i])
		if err != nil {
			errs = append(errs, err)
		}
	}

	for i := range ipFiles {
		var bs []byte
		bs, err = os.ReadFile(ipFiles[i])
		if err != nil {
			errs = append(errs, err)
			continue
		}
		bIPs := bytes.Fields(bs)
		for j := range bIPs {
			err = man.AddByString(string(bIPs[j]))
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if !man.HasData() {
		man = nil
	}

	return
}

func getIPAllowMiddleware(ips, ipFiles []string, outputMids []middleware.Middleware) (middleware.Middleware, []error) {
	man, errs := getIPRangeMan(ips, ipFiles)
	if man == nil || len(errs) > 0 {
		return nil, errs
	}

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		ip, _ := util.ExtractIPPort(r.RemoteAddr)
		if man.MatchStringAddr(ip) {
			return middleware.GoNext
		}

		util.LogErrorString(context.Logger, "request denied as out of allow list from "+r.RemoteAddr)
		result = middleware.Outputted
		context.Status = http.StatusForbidden
		for i := range outputMids {
			if outputMids[i](w, r, context) == middleware.Outputted {
				return
			}
		}

		w.WriteHeader(context.Status)
		return
	}, nil
}

func getIPDenyMiddleware(ips, ipFiles []string, outputMids []middleware.Middleware) (middleware.Middleware, []error) {
	man, errs := getIPRangeMan(ips, ipFiles)
	if man == nil || len(errs) > 0 {
		return nil, errs
	}

	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		ip, _ := util.ExtractIPPort(r.RemoteAddr)
		if !man.MatchStringAddr(ip) {
			return middleware.GoNext

		}

		util.LogErrorString(context.Logger, "request denied as match deny list from "+r.RemoteAddr)
		result = middleware.Outputted
		context.Status = http.StatusForbidden
		for i := range outputMids {
			if outputMids[i](w, r, context) == middleware.Outputted {
				return
			}
		}

		w.WriteHeader(context.Status)
		return
	}, nil
}
