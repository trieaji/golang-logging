package golangloging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct { //untuk mengimplementasikan interface hook

}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel} //yg ke trigger levelnya adalah error dan warn
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("sample hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Hello info")
	logger.Warn("Hello warn")
	logger.Error("Hello error")
}

//Hook itu semacam callback atau listener ketika sebuah kejadian terjadi