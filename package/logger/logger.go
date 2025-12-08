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

func newConfig(logConfig *LogConfig) *LogConfig {
	if logConfig.Writer == nil {
		logConfig.Writer = os.Stdout
	}
	if logConfig.LogLevel == "" {
		logConfig.LogLevel = DEFAULT_LOG_LEVEL
	}
	logConfig.LogLevel = strings.ToUpper(logConfig.LogLevel)
	return logConfig
}

func New(logConfig LogConfig) *LambdaLogger {
	config := newConfig(&logConfig)
	zerolog.SetGlobalLevel(LogMapper[config.LogLevel])
	zerolog.CallerFieldName = CALLER_NAME
	return &LambdaLogger{
		logger: zerolog.
			New(logConfig.Writer).
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
