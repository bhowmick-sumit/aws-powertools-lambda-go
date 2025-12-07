package logger

import (
	"os"

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
	if logConfig.logLevel == "" {
		logConfig.logLevel = DEFAULT_LOG_LEVEL
	}
	return &logConfig
}

func New(logConfig LogConfig) *LambdaLogger {
	config := NewConfig(logConfig)
	zerolog.SetGlobalLevel(LogMapper[config.logLevel])
	zerolog.CallerFieldName = CALLER_NAME
	return &LambdaLogger{
		logger: zerolog.
			New(os.Stdout).
			With().
			CallerWithSkipFrameCount(CALLER_SKIP_FRAME_COUNT).
			Timestamp().
			Logger(),
	}
}

func (log *LambdaLogger) Fatal(message string) {
	log.logger.Fatal().Msg(message)
}

func (log *LambdaLogger) Error(message string) {
	log.logger.Error().Msg(message)
}

func (log *LambdaLogger) Warn(message string) {
	log.logger.Warn().Msg(message)
}

func (log *LambdaLogger) Info(message string) {
	log.logger.Info().Msg(message)
}

func (log *LambdaLogger) Debug(message string) {
	log.logger.
		Debug().
		Msg(message)
}

func (log *LambdaLogger) Trace(message string) {
	log.logger.Trace().Msg(message)
}
