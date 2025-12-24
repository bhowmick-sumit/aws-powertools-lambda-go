package logger

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/suctl/aws-powertools-lambda-go/logger/types"
)

func TestLoggerNewConfigWithDefaults(t *testing.T) {
	config := newConfig(&types.LogConfig{})
	if config.Writer != os.Stdout {
		t.Errorf("expected default writer to be os.Stdout")
	}
}

func TestLoggerNewConfigWithCustomProperties(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	config := newConfig(&types.LogConfig{
		Writer: os.Stderr,
	})
	if config.Writer != os.Stderr {
		t.Errorf("expected default writer to be os.Stdout")
	}
}

func TestLoggerWithDefaultConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	logger := New(types.LogConfig{})
	logger.Info("Info Log")
	logger.Debug("Debug Log")
	logger.Warn("Warn Log")
	logger.Error("Error Log")
}

func TestLoggerWithCustomConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	logger := New(types.LogConfig{})
	logger.Info("info log")
	logger.Debug("debug log")
	logger.Trace("trace log")
	logger.Warn("warn log")
}

func TestLoggerWithLogLevelInfo(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Info("info log")
	logger.Debug("debug log")

	output := buf.String()
	if strings.Contains(output, "debug log") {
		t.Errorf("debug log should not be present")
	}
}

func TestLoggerWithErrorLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Error("error log")
	output := buf.String()
	if !strings.Contains(output, "error log") {
		t.Errorf("expected output to contain 'error log', got '%s'", output)
	}
	if !strings.Contains(output, "\"level\":\"error\"") {
		t.Errorf("expected output to contain 'level': 'error', got '%s'", output)
	}
}

func TestLoggerWithWarnLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "WARN")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Warn("warn log")
	output := buf.String()
	if !strings.Contains(output, "warn log") {
		t.Errorf("expected output to contain 'warn log', got '%s'", output)
	}
	if !strings.Contains(output, "\"level\":\"warn\"") {
		t.Errorf("expected output to contain 'level': 'warn', got '%s'", output)
	}
}

func TestLoggerWithInfoLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Info("info log")
	output := buf.String()
	if !strings.Contains(output, "info log") {
		t.Errorf("expected output to contain 'info log', got '%s'", output)
	}
	if !strings.Contains(output, "\"level\":\"info\"") {
		t.Errorf("expected output to contain 'level': 'info', got '%s'", output)
	}
}

func TestLoggerWithDebugLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "DEBUG")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Debug("debug log")
	output := buf.String()
	if !strings.Contains(output, "debug log") {
		t.Errorf("expected output to contain 'debug log', got '%s'", output)
	}
	if !strings.Contains(output, "\"level\":\"debug\"") {
		t.Errorf("expected output to contain 'level': 'debug', got '%s'", output)
	}
}

func TestLoggerWithTraceLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "TRACE")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
	})
	logger.Trace("trace log")
	output := buf.String()
	if !strings.Contains(output, "trace log") {
		t.Errorf("expected output to contain 'trace log', got '%s'", output)
	}
	if !strings.Contains(output, "\"level\":\"trace\"") {
		t.Errorf("expected output to contain 'level': 'trace', got '%s'", output)
	}
}

func TestLoggerWithProperties(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	defer func() {
		os.Unsetenv("POWERTOOLS_LOG_LEVEL")
	}()
	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
		Properties: map[string]string{
			"name": "sumit",
			"env":  "production",
		},
	})
	logger.Info("info log with properties")
	output := buf.String()
	if !strings.Contains(output, "\"name\":\"sumit\"") {
		t.Errorf("expected output to contain 'name': 'sumit', got '%s'", output)
	}
	if !strings.Contains(output, "\"env\":\"production\"") {
		t.Errorf("expected output to contain 'env': 'production', got '%s'", output)
	}
}

func TestLoggerInjectContext(t *testing.T) {
	os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "test-function")
	os.Setenv("AWS_LAMBDA_FUNCTION_INVOKED_ARN", "arn:aws:lambda:us-east-1:123456789012:function:test-function")
	os.Setenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE", "128")

	defer func() {
		os.Unsetenv("AWS_LAMBDA_FUNCTION_NAME")
		os.Unsetenv("AWS_LAMBDA_FUNCTION_INVOKED_ARN")
		os.Unsetenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE")
	}()

	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
		Properties: map[string]string{
			"service": "test-service",
		},
	})

	ctx := context.WithValue(context.Background(), "aws_request_id", "test-request-id")
	ctx = context.WithValue(ctx, "invoked_function_arn", "arn:aws:lambda:us-east-1:123456789012:function:test-function")

	logger.InjectContext(ctx)
	logger.Info("info log with lambda context")

	output := buf.String()
	if !strings.Contains(output, "\"function_name\":\"test-function\"") {
		t.Errorf("expected output to contain 'function_name': 'test-function', got '%s'", output)
	}
	if !strings.Contains(output, "\"function_memory_size\":\"128\"") {
		t.Errorf("expected output to contain 'function_memory_size': '128', got '%s'", output)
	}
	if !strings.Contains(output, "\"function_arn\":\"arn:aws:lambda:us-east-1:123456789012:function:test-function\"") {
		t.Errorf("expected output to contain correct 'function_arn', got '%s'", output)
	}
	if !strings.Contains(output, "\"function_request_id\":\"test-request-id\"") {
		t.Errorf("expected output to contain 'function_request_id': 'test-request-id', got '%s'", output)
	}
	if !strings.Contains(output, "\"service\":\"test-service\"") {
		t.Errorf("expected output to contain 'service': 'test-service', got '%s'", output)
	}
}

func TestLoggerInjectContextWhenFailed(t *testing.T) {
	defer func() {
		os.Unsetenv("AWS_LAMBDA_FUNCTION_NAME")
		os.Unsetenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE")
	}()

	var buf bytes.Buffer
	logger := New(types.LogConfig{
		Writer: &buf,
		Properties: map[string]string{
			"service": "test-service",
		},
	})

	ctx := context.Background()

	logger.InjectContext(ctx)
	logger.Info("info log with lambda context")

	output := buf.String()
	if !strings.Contains(output, "\"function_request_id\":\"unknown\"") {
		t.Errorf("expected output to contain 'function_request_id': 'unknown', got '%s'", output)
	}
	if !strings.Contains(output, "\"function_name\":\"unknown\"") {
		t.Errorf("expected output to contain 'function_name': 'unknown', got '%s'", output)
	}
	if !strings.Contains(output, "\"function_memory_size\":\"0\"") {
		t.Errorf("expected output to contain 'function_memory_size': '0', got '%s'", output)
	}
	if !strings.Contains(output, "\"service\":\"test-service\"") {
		t.Errorf("expected output to contain 'service': 'test-service', got '%s'", output)
	}
}
