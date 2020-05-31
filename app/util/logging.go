package util

import (
	"fmt"

	"logur.dev/logur"
)

func LogDebug(logger logur.Logger, msg string, params ...interface{}) {
	logger.Info(fmt.Sprintf(msg, params...))
}

func LogInfo(logger logur.Logger, msg string, params ...interface{}) {
	logger.Info(fmt.Sprintf(msg, params...))
}

func LogWarn(logger logur.Logger, msg string, params ...interface{}) {
	logger.Warn(fmt.Sprintf(msg, params...))
}

func LogError(logger logur.Logger, msg string, params ...interface{}) {
	logger.Error(fmt.Sprintf(msg, params...))
}
