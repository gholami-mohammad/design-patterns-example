package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type ZerologAdapter struct {
}

func (l *ZerologAdapter) Info(msg ...interface{}) {
	log := zerolog.New(os.Stdout)
	str := ""
	for m := range msg {
		str += fmt.Sprintf("%v; ", m)
	}
	log.Info().Msg(str)
}

func (l *ZerologAdapter) Error(err ...interface{}) {
	log := zerolog.New(os.Stdout)
	str := ""
	for m := range err {
		str += fmt.Sprintf("%v; ", m)
	}
	log.Error().Msg(str)
}
