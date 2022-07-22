package main

import (
	"adapter/logger"
	"errors"
)

func main() {
	l := logger.NewLogger(new(logger.LogrusAdapter))
	l.Info("Info logged by logrus")
	l.Error("Error1", errors.New("Error2"))

	l2 := logger.NewLogger(new(logger.ZerologAdapter))
	l2.Info("Info1", "info2")
	l2.Error(errors.New("E1"), "E2")
}
