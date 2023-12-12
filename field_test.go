package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Laksa").Info("Hello Laksa")

	//field lebih dari satu
	logger.WithField("username", "Laksa").WithField("name", "Bayu").Info("Hello Laksa dan Bayu")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "laksa",
		"name": "Laksa Bayu", 
	}).Info("Hello world")
}