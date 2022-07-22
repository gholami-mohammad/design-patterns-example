package logger

import "github.com/sirupsen/logrus"

type LogrusAdapter struct {
}

func (l *LogrusAdapter) Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func (l *LogrusAdapter) Error(err ...interface{}) {
	logrus.Error(err...)
}
