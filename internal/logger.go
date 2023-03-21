package internal

import (
	"os"

	"bitbucket.org/libertywireless/circles-sandbox/domain"
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
