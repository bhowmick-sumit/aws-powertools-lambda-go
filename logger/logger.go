package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/suctl/aws-powertools-lambda-go/internal/utils"
	"github.com/suctl/aws-powertools-lambda-go/logger/types"
)

const (
	callerName           = "location"
	callerSkipFrameCount = 3
	defaultLogLevel      = "DEBUG"
)

var LogMapper = map[string]zerolog.Level{
	"ERROR": zerolog.ErrorLevel,
	"WARN":  zerolog.WarnLevel,
	"INFO":  zerolog.InfoLevel,
	"DEBUG": zerolog.DebugLevel,
	"TRACE": zerolog.TraceLevel,
}

type Logger struct {
	logger zerolog.Logger
}

func (log *Logger) Error(message string, args ...any) {
	log.logger.Error().Msgf(message, args...)
}

func (log *Logger) Warn(message string, args ...any) {
	log.logger.Warn().Msgf(message, args...)
}

func (log *Logger) Info(message string, args ...any) {
	log.logger.Info().Msgf(message, args...)
}

func (log *Logger) Debug(message string, args ...any) {
	log.logger.Debug().Msgf(message, args...)
}

func (log *Logger) Trace(message string, args ...any) {
	log.logger.Trace().Msgf(message, args...)
}

func New(logConfig types.LogConfig) *Logger {
	setConfigFromEnvironment()
	config := newConfig(&logConfig)
	zerolog.CallerFieldName = callerName
	logger := zerolog.
		New(config.Writer).
		With().
		CallerWithSkipFrameCount(callerSkipFrameCount).
		Timestamp().
		Logger()

	for key, value := range logConfig.Properties {
		logger = logger.With().Str(key, value).Logger()
	}

	return &Logger{
		logger: logger,
	}
}

func setConfigFromEnvironment() {
	logLevel := utils.GetEnvironmentVariable("POWERTOOLS_LOG_LEVEL", defaultLogLevel)
	zerolog.SetGlobalLevel(LogMapper[strings.ToUpper(logLevel)])
}

func newConfig(logConfig *types.LogConfig) *types.LogConfig {
	if logConfig.Writer == nil {
		logConfig.Writer = os.Stdout
	}
	return logConfig
}
