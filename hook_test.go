package go_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()

	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello info")
	logger.Warn("Hello warn")
	logger.Error("Hello error")
	logger.Debug("Hello warn")
}
