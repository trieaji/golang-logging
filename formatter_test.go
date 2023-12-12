package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello info")
	logger.Warn("Hello warn")
	logger.Error("Hello error")
}