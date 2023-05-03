package mylogger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func NewLogger() *zerolog.Logger {
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		//Int("pid", os.Getpid()).
		//Str("go_version", buildInfo.GoVersion).
		Logger()
	return &l
}
