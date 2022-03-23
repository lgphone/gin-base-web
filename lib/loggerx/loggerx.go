package loggerx

import (
	"github.com/sirupsen/logrus"
	"os"
	"younghe/config"
)

var Logger *logrus.Logger

func Setup() {
	Logger = logrus.StandardLogger()
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(true)
	if config.Config.Debug {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
}
