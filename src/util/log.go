package util

import "mjpclab.dev/ghfs/src/serverLog"

func LogAccess(logger *serverLog.Logger, msg string) {
	if logger == nil {
		return
	}

	if logger.CanLogAccess() {
		buf := serverLog.NewBuffer(len(msg))
		buf = append(buf, []byte(msg)...)
		logger.LogAccess(buf)
	}
}

func LogError(logger *serverLog.Logger, err error) {
	if logger == nil {
		return
	}

	if logger.CanLogError() {
		strErr := err.Error()
		buf := serverLog.NewBuffer(len(strErr))
		buf = append(buf, []byte(strErr)...)
		logger.LogError(buf)
	}
}

func LogErrorString(logger *serverLog.Logger, err string) {
	if logger == nil {
		return
	}

	if logger.CanLogError() {
		buf := serverLog.NewBuffer(len(err))
		buf = append(buf, []byte(err)...)
		logger.LogError(buf)
	}
}
