package logger

import (
	"os"
	"strings"

	"github.com/bhowmick-sumit/aws-powertools-lambda-go/internal/utils"
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

func setConfigFromEnvironment() {
	logLevel := utils.GetEnvironmentVariable("POWERTOOLS_LOG_LEVEL", DEFAULT_LOG_LEVEL)
	zerolog.SetGlobalLevel(LogMapper[strings.ToUpper(logLevel)])
}

func newConfig(logConfig *LogConfig) *LogConfig {
	if logConfig.Writer == nil {
		logConfig.Writer = os.Stdout
	}
	return logConfig
}

func New(logConfig LogConfig) *LambdaLogger {
	setConfigFromEnvironment()
	config := newConfig(&logConfig)
	zerolog.CallerFieldName = CALLER_NAME
	logger := zerolog.
		New(config.Writer).
		With().
		CallerWithSkipFrameCount(CALLER_SKIP_FRAME_COUNT).
		Timestamp().
		Logger()

	for key, value := range logConfig.Properties {
		logger = logger.With().Str(key, value).Logger()
	}

	return &LambdaLogger{
		logger: logger,
	}
}

// func (log *LambdaLogger) addPropertie(key, value string) zerolog.Logger {
// 	return log.logger.With().Str(key, value).Logger()
// }

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
