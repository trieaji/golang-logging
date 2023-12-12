package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello info")
	logrus.Warn("Hello warn")

	//mengubah logrus singleton yg global menjadi formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello info")
	logrus.Warn("Hello warn")
}