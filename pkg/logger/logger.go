// Package logger is describe reusable package for logger app
package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

func NewLogger(formatter logrus.Formatter) *Logger {
	log := logrus.New()

	log.SetFormatter(formatter)

	return &Logger{
		Logger: log,
	}
}