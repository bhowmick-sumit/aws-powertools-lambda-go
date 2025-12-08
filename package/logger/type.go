package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type LogConfig struct {
	LogLevel string
	Writer   io.Writer
}

type LambdaLogger struct {
	logger zerolog.Logger
}
