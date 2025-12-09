package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type LogConfig struct {
	Writer     io.Writer
	Properties map[string]string
}

type LambdaLogger struct {
	logger zerolog.Logger
}
