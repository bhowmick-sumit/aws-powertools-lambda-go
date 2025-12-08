package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	CALLER_NAME             = "location"
	CALLER_SKIP_FRAME_COUNT = 3
	DEFAULT_LOG_LEVEL       = "DEBUG"
)

var LogMapper = map[string]zerolog.Level{
	"FATAL": zerolog.FatalLevel,
	"ERROR": zerolog.ErrorLevel,
	"WARN":  zerolog.WarnLevel,
	"INFO":  zerolog.InfoLevel,
	"DEBUG": zerolog.DebugLevel,
	"TRACE": zerolog.TraceLevel,
}

func NewConfig(logConfig LogConfig) *LogConfig {
	if logConfig.writer == nil {
		logConfig.writer = os.Stdout
	}
	if logConfig.logLevel == "" {
		logConfig.logLevel = DEFAULT_LOG_LEVEL
	}
	logConfig.logLevel = strings.ToUpper(logConfig.logLevel)
	return &logConfig
}

func New(logConfig LogConfig) *LambdaLogger {
	config := NewConfig(logConfig)
	zerolog.SetGlobalLevel(LogMapper[config.logLevel])
	zerolog.CallerFieldName = CALLER_NAME
	return &LambdaLogger{
		logger: zerolog.
			New(logConfig.writer).
			With().
			CallerWithSkipFrameCount(CALLER_SKIP_FRAME_COUNT).
			Timestamp().
			Logger(),
	}
}

func (log *LambdaLogger) Error(message string, args ...any) {
	log.logger.Error().Msgf(message, args...)
}

func (log *LambdaLogger) Warn(message string, args ...any) {
	log.logger.Warn().Msgf(message, args...)
}

func (log *LambdaLogger) Info(message string, args ...any) {
	log.logger.Info().Msgf(message, args...)
}

func (log *LambdaLogger) Debug(message string, args ...any) {
	log.logger.Debug().Msgf(message, args...)
}

func (log *LambdaLogger) Trace(message string, args ...any) {
	log.logger.Trace().Msgf(message, args...)
}
