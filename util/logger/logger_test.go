package logger

import (
	"bytes"
	"regexp"
	"testing"
)

func TestLogger(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new Logger with DebugLevel
	logger := New(&buf, DebugLevel)

	// Test Debug
	logger.Debug("Debug message")
	if buf.String() != "Debug message" {
		t.Errorf("Expected 'Debug message', got '%s'", buf.String())
	}
	buf.Reset()

	// Test Debugln
	logger.Debugln("Debug line")
	if buf.String() != "Debug line\n" {
		t.Errorf("Expected 'Debug line\n', got '%s'", buf.String())
	}
	buf.Reset()

	// Test Debugf
	logger.Debugf("Debug formatted %s", "message")
	if buf.String() != "Debug formatted message" {
		t.Errorf("Expected 'Debug formatted message', got '%s'", buf.String())
	}
	buf.Reset()

	// Test Warning
	logger.Warning("Warning message")
	if buf.String() != "Warning message" {
		t.Errorf("Expected 'Warning message', got '%s'", buf.String())
	}
	buf.Reset()

	// Test Warningln
	logger.Warningln("Warning line")
	if buf.String() != "Warning line\n" {
		t.Errorf("Expected 'Warning line\n', got '%s'", buf.String())
	}
	buf.Reset()

	// Test Warningf
	logger.Warningf("Warning formatted %s", "message")
	if buf.String() != "Warning formatted message" {
		t.Errorf("Expected 'Warning formatted message', got '%s'", buf.String())
	}
	buf.Reset()
}

func TestLoggerHelpers(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new Logger with DebugLevel
	logger := New(&buf, DebugLevel)

	// Test SetOutput
	newBuf := bytes.Buffer{}
	logger.SetOutput(&newBuf)
	logger.Debug("Debug message with new output")
	if newBuf.String() != "Debug message with new output" {
		t.Errorf("Expected 'Debug message with new output', got '%s'", newBuf.String())
	}

	// Test SetLevel and Level
	logger.SetLevel(WarningLevel)
	if logger.Level() != WarningLevel {
		t.Errorf("Expected log level %d, got %d", WarningLevel, logger.Level())
	}

	// Test IsDebugLevel
	if logger.IsDebugLevel() {
		t.Errorf("Expected IsDebugLevel() to be false")
	}

	// Test IsWarningLevel
	if !logger.IsWarningLevel() {
		t.Errorf("Expected IsWarningLevel() to be true")
	}

	// Test Default
	defaultLogger := Default()
	if defaultLogger != std {
		t.Errorf("Expected Default() to return global instance")
	}
}

func TestNewLogger(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new logger using NewLogger
	logger := NewLogger("Prefix: ")

	// Set the output to the buffer
	logger.SetOutput(&buf)

	// Log a message
	logger.Println("Log message")

	// Check if the buffer contains the log message
	expected := "Prefix: Log message\n"
	actual := buf.String()

	// Define a regular expression pattern to match the log message without the timestamp
	pattern := `Prefix: \d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} Log message\n`

	// Use regex matching to compare the actual output with the pattern
	matched, err := regexp.MatchString(pattern, actual)
	if err != nil {
		t.Errorf("Error while matching regex pattern: %v", err)
	}

	if !matched {
		t.Errorf("Expected output did not match the pattern:\nExpected: '%s'\nActual: '%s'", expected, actual)
	}
}

func TestLoggerGlobalMethods(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Set the output to the buffer
	SetOutput(&buf)

	// Test IsDebugLevel
	if IsDebugLevel() {
		t.Errorf("Expected IsDebugLevel() to be false")
	}

	// Test IsWarningLevel
	if !IsWarningLevel() {
		t.Errorf("Expected IsWarningLevel() to be true")
	}

	t.Run("Test Debug", func(t *testing.T) {
		_, err := Debug("Global debug message")
		if err != nil {
			t.Errorf("Error while logging debug message: %v", err)
		}
	})

	// Test Debugln
	t.Run("Test Debugln", func(t *testing.T) {
		_, err := Debugln("Global debug line")
		if err != nil {
			t.Errorf("Error while logging debug line: %v", err)
		}
	})

	// Test Debugf
	t.Run("Test Debugf", func(t *testing.T) {
		_, err := Debugf("Global debug formatted %s", "message")
		if err != nil {
			t.Errorf("Error while logging formatted debug message: %v", err)
		}
	})
	// Test Warn
	Warn("Global warning message")
	expected := "Global warning message"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
	buf.Reset()

	// Test Warnln
	Warnln("Global warning line")
	expected = "Global warning line\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
	buf.Reset()

	// Test Warnf
	Warnf("Global warning formatted %s", "message")
	expected = "Global warning formatted message"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}
