package src

import (
	"errors"
	"mjpclab.dev/ehfs/src/middleware"
	"mjpclab.dev/ehfs/src/param"
	"mjpclab.dev/ehfs/src/version"
	"mjpclab.dev/ghfs/src/app"
	"mjpclab.dev/ghfs/src/serverError"
	"mjpclab.dev/ghfs/src/setting"
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
		os.Exit(0)
	}()
}

func reopenLogOnHup(appInst *app.App) {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGHUP)

	go func() {
		for range chSignal {
			errs := appInst.ReOpenLog()
			serverError.CheckFatal(errs...)
		}
	}()
}

func Main() {
	// params
	baseParams, params, printVersion, printHelp, errs := param.ParseFromCli()
	serverError.CheckFatal(errs...)
	if printVersion {
		version.PrintVersion()
		os.Exit(0)
	}
	if printHelp {
		param.PrintHelp()
		os.Exit(0)
	}

	// apply middlewares
	errs = middleware.ApplyMiddlewares(baseParams, params)
	serverError.CheckFatal(errs...)

	// setting
	setting := setting.ParseFromEnv()

	// app
	appInst, errs := app.NewApp(baseParams, setting)
	serverError.CheckFatal(errs...)
	if appInst == nil {
		serverError.CheckFatal(errors.New("failed to create application instance"))
	}

	cleanupOnEnd(appInst)
	reopenLogOnHup(appInst)
	errs = appInst.Open()
	serverError.CheckFatal(errs...)
	appInst.Shutdown()
}
