package src

import (
	"mjpclab.dev/ehfs/src/middleware"
	"mjpclab.dev/ehfs/src/param"
	"mjpclab.dev/ehfs/src/version"
	"mjpclab.dev/ghfs/src"
	"mjpclab.dev/ghfs/src/serverError"
	"mjpclab.dev/ghfs/src/setting"
)

func Main() (ok bool) {
	// params
	baseParams, params, printVersion, printHelp, errs := param.ParseFromCli()
	if serverError.CheckError(errs...) {
		return
	}
	if printVersion {
		version.PrintVersion()
		return true
	}
	if printHelp {
		param.PrintHelp()
		return true
	}

	// apply middlewares
	errs = middleware.ApplyMiddlewares(baseParams, params)
	if serverError.CheckError(errs...) {
		return
	}

	// settings
	settings := setting.ParseFromEnv()

	// start
	errs = src.Start(settings, baseParams)
	if serverError.CheckError(errs...) {
		return
	}

	return true
}
