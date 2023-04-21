package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "taufiq").
		Info("Hello World")
	logger.WithField("username", "kurniawan").
		WithField("name", "bayu").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "taufiq",
		"name":     "kurniawan",
	}).Info("Hello World")
}
