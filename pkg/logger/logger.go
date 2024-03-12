package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	var logs = logrus.New()

	logs.Formatter = new(logrus.JSONFormatter)
	logs.Level = logrus.InfoLevel
	logs.Out = os.Stdout

	return logs
}
