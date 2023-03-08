package logService

import logrusInfra "backend-skeleton-golang/commons/infra/logrus"

func Info(msg string) {
	logrusInfra.Log.Info(msg)
}

func Error(msg string) {
	logrusInfra.Log.Error(msg)
}
