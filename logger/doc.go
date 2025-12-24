// Package logger provides a structured logging utility built on top of
// zerolog, designed specifically for AWS Lambda workloads.
//
// The logger automatically:
//   - Configures log level from environment variables
//   - Emits JSON structured logs
//   - Enriches logs with AWS Lambda context (request ID, ARN, memory, etc.)
//   - Supports contextual and formatted logging
//
// # Log Level Configuration
//
// The log level is controlled using the environment variable:
//
//	POWERTOOLS_LOG_LEVEL
//
// Supported values:
//
//	ERROR, WARN, INFO, DEBUG, TRACE
//
// If the variable is not set, the default log level is DEBUG.
//
// # AWS Lambda Context Injection
//
// When running inside AWS Lambda, the logger can inject runtime metadata
// from the Lambda execution context:
//
//   - function_name
//   - function_memory_size
//   - function_arn
//   - function_request_id
//
// This is achieved by calling InjectContext with the Lambda context.
//
// # Usage
//
// Basic usage:
//
//	log := logger.New(types.LogConfig{})
//	log.Info("service started")
//
// Injecting Lambda context:
//
//	func handler(ctx context.Context) {
//	    log := logger.New(types.LogConfig{})
//	    log.InjectContext(ctx)
//	    log.Info("request received")
//	}
//
// # Structured Properties
//
// Static key-value properties can be attached to every log entry using
// LogConfig.Properties:
//
//	log := logger.New(types.LogConfig{
//	    Properties: map[string]string{
//	        "service": "orders",
//	        "env":     "prod",
//	    },
//	})
//
// # Output
//
// Logs are emitted in JSON format and written to stdout by default,
// making them compatible with AWS CloudWatch Logs and log aggregators.
//
// This package is inspired by AWS Powertools logging patterns and
// follows best practices for observability in serverless environments.
package logger
