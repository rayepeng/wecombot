package wecom

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	configureLogger(Logger)
	logMode := os.Getenv("WECOM_LOG_MODE")
	if logMode == "stdout" {
		Logger.SetOutput(os.Stdout)
	} else {
		logFile, err := os.OpenFile("wecom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Logger.SetOutput(logFile)
		} else {
			Logger.SetOutput(os.Stdout)
			Logger.WithError(err).Error("Failed to open log file, using default stderr")
		}
	}

}

func configureLogger(logger *logrus.Logger) {
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		CallerPrettyfier: callerPrettyfier,
	})
	logger.SetReportCaller(true)
}

func callerPrettyfier(f *runtime.Frame) (string, string) {
	filename := path.Base(f.File)
	return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
}
