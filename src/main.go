package src

import (
	"errors"
	localDefaultTheme "mjpclab.dev/ehfs/src/defaultTheme"
	"mjpclab.dev/ehfs/src/middleware"
	"mjpclab.dev/ehfs/src/param"
	"mjpclab.dev/ehfs/src/version"
	"mjpclab.dev/ghfs/src"
	"mjpclab.dev/ghfs/src/app"
	"mjpclab.dev/ghfs/src/serverError"
	"mjpclab.dev/ghfs/src/setting"
	"mjpclab.dev/ghfs/src/tpl/defaultTheme"
	"os"
	"os/signal"
	"syscall"
)

func cleanupOnEnd(appInst *app.App) {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-chSignal
		appInst.Shutdown()
	}()
}

func reInitOnHup(appInst *app.App) {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGHUP)

	go func() {
		for range chSignal {
			errs := appInst.ReOpenLog()
			if serverError.CheckError(errs...) {
				appInst.Shutdown()
				break
			}
			errs = appInst.ReLoadCertificates()
			if serverError.CheckError(errs...) {
				appInst.Shutdown()
				break
			}
		}
	}()
}

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

	// override default theme
	defaultTheme.DefaultTheme = localDefaultTheme.DefaultTheme

	// settings
	settings := setting.ParseFromEnv()

	// CPU profile
	if len(settings.CPUProfileFile) > 0 {
		cpuProfileFile, err := src.StartCPUProfile(settings.CPUProfileFile)
		if serverError.CheckError(err) {
			return
		}
		defer src.StopCPUProfile(cpuProfileFile)
	}

	// app
	appInst, errs := app.NewApp(baseParams, settings)
	if serverError.CheckError(errs...) {
		return
	}
	if appInst == nil {
		serverError.CheckError(errors.New("failed to create application instance"))
		return
	}

	cleanupOnEnd(appInst)
	reInitOnHup(appInst)
	errs = appInst.Open()
	if serverError.CheckError(errs...) {
		return
	}

	return true
}
