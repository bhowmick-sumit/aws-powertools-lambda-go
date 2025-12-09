package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestLoggerNewConfigWithDefaults(t *testing.T) {
	config := newConfig(&LogConfig{})
	if config.Writer != os.Stdout {
		t.Errorf("expected default writer to be os.Stdout")
	}
}

func TestLoggerNewConfigWithCustomProperties(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	config := newConfig(&LogConfig{
		Writer: os.Stderr,
	})
	if config.Writer != os.Stderr {
		t.Errorf("expected default writer to be os.Stdout")
	}
}

func TestLambdaLoggerWithDefaultConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	logger := New(LogConfig{})
	logger.Info("Info Log")
	logger.Debug("Debug Log")
	logger.Warn("Warn Log")
	logger.Error("Error Log")
}

func TestLambdaLoggerWithCustomConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	logger := New(LogConfig{})
	logger.Info("info log")
	logger.Debug("debug log")
	logger.Trace("trace log")
	logger.Warn("warn log")
}

func TestLambdaLoggerWithLogLevelInfo(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	var buf bytes.Buffer
	logger := New(LogConfig{
		Writer: &buf,
	})
	logger.Info("info log")
	logger.Debug("debug log")

	output := buf.String()
	if strings.Contains(output, "debug log") {
		t.Errorf("debug log should not be present")
	}
}

func TestLambdaLoggerWithErrorLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	var buf bytes.Buffer
	logger := New(LogConfig{
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

func TestLambdaLoggerWithWarnLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "WARN")
	var buf bytes.Buffer
	logger := New(LogConfig{
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

func TestLambdaLoggerWithInfoLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	var buf bytes.Buffer
	logger := New(LogConfig{
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

func TestLambdaLoggerWithDebugLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "DEBUG")
	var buf bytes.Buffer
	logger := New(LogConfig{
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

func TestLambdaLoggerWithTraceLog(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "TRACE")
	var buf bytes.Buffer
	logger := New(LogConfig{
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

func TestLambdaLoggerWithProperties(t *testing.T) {
	os.Setenv("POWERTOOLS_LOG_LEVEL", "INFO")
	var buf bytes.Buffer
	logger := New(LogConfig{
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
