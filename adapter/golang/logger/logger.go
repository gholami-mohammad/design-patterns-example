package logger

type LoggerAdapter interface {
	Info(msg ...interface{})
	Error(err ...interface{})
}

func NewLogger(loggeraAdapter LoggerAdapter) logger {
	return logger{loggerAdapter: loggeraAdapter}
}

type logger struct {
	loggerAdapter LoggerAdapter
}

func (l *logger) Info(msg ...interface{}) {
	l.loggerAdapter.Info(msg...)
}

func (l *logger) Error(err ...interface{}) {
	l.loggerAdapter.Error(err...)
}
