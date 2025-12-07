package logger

import "github.com/rs/zerolog"

type LogConfig struct {
	logLevel       string
	inject_context bool
}

type LambdaLogger struct {
	logger zerolog.Logger
}
