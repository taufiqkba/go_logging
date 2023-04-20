package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}
