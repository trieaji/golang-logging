package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

//dibahasa karena detailnya itu sebenarnya ada object yg bernama entry
func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logging")
	logger.WithField("username", "bayu").Info("Hello broo")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "laksa")
	entry.Info("Hello entry")
}