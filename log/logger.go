package log

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerSetup sync.Once

// SetupLogger creates logger instance
func SetupLogger() {
	loggerSetup.Do(func() {
		logger = logrus.New()
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
}

//GetLogger returns logger for logging generic logs
func GetLogger() *logrus.Logger {
	return logger
}
