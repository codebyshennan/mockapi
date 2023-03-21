package internal

import (
	"os"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/sirupsen/logrus"
)

func getLogger() domain.ILogger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}
