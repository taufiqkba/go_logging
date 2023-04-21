package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "kurniawan").Info("Hello Logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "taufiq")
	entry.Info("Hello Entry")

}
