package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type LogConfig struct {
	logLevel       string
	writer         io.Writer
	inject_context bool
}

type LambdaLogger struct {
	logger zerolog.Logger
}
