package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

const (
	NoneLevel Level = iota
	WarningLevel
	DebugLevel
)

type Logger struct {
	w     io.Writer
	level Level
}

// New creates a new Logger object with the given writer and log level
func New(w io.Writer, l Level) *Logger {
	// Create a new Logger object with the given writer and log level
	return &Logger{
		level: l,
		w:     w,
	}
}

// NewLogger sets up a logger that writes to the stderr
func NewLogger(prefix string) *log.Logger {
	// we create our own stderr since we're going to nuke the existing one
	return log.New(os.Stderr, prefix, log.LstdFlags)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.w = w
}

// SetLevel sets the log level for the Logger object
func (l *Logger) SetLevel(level Level) {
	// Set the level field of the Logger object to the given level
	l.level = level
}

// Level returns the log level of the Logger object
func (l *Logger) Level() Level {
	// Return the level field of the Logger object
	return l.level
}

// IsDebugLevel returns true if the log level of the Logger object is DebugLevel or higher
func (l *Logger) IsDebugLevel() bool {
	// If the level field of the Logger object is greater than or equal to DebugLevel, return true
	return l.level >= DebugLevel
}

// IsWarningLevel returns true if the log level of the Logger object is WarningLevel or higher
func (l *Logger) IsWarningLevel() bool {
	// If the level field of the Logger object is greater than or equal to WarningLevel, return true
	return l.level >= WarningLevel
}

// Debug writes the given arguments to the Logger object if the log level is DebugLevel or higher
func (l Logger) Debug(a ...any) (int, error) {
	// If the log level is not DebugLevel or higher, return 0 and nil
	if !l.IsDebugLevel() {
		return 0, nil
	}

	// Write the given arguments to the Logger object and return the number of bytes written and any error
	return fmt.Fprint(l.w, a...)
}

// Debugln writes the given arguments to the Logger object with a newline character if the log level is DebugLevel or higher
func (l Logger) Debugln(a ...any) (int, error) {
	// If the log level is not DebugLevel or higher, return 0 and nil
	if !l.IsDebugLevel() {
		return 0, nil
	}

	// Write the given arguments to the Logger object with a newline character and return the number of bytes written and any error
	return fmt.Fprintln(l.w, a...)
}

// Debugf writes the given formatted string and arguments to the Logger object if the log level is DebugLevel or higher
func (l Logger) Debugf(format string, a ...any) (int, error) {
	// If the log level is not DebugLevel or higher, return 0 and nil
	if !l.IsDebugLevel() {
		return 0, nil
	}

	// Write the given formatted string and arguments to the Logger object and return the number of bytes written and any error
	return fmt.Fprintf(l.w, format, a...)
}

// Warning writes the given arguments to the Logger object if the log level is WarningLevel or higher
func (l Logger) Warning(a ...any) (int, error) {
	// If the log level is not WarningLevel or higher, return 0 and nil
	if !l.IsWarningLevel() {
		return 0, nil
	}

	// Write the given arguments to the Logger object and return the number of bytes written and any error
	return fmt.Fprint(l.w, a...)
}

// Warningln writes the given arguments to the Logger object with a newline character if the log level is WarningLevel or higher
func (l Logger) Warningln(a ...any) (int, error) {
	// If the log level is not WarningLevel or higher, return 0 and nil
	if !l.IsWarningLevel() {
		return 0, nil
	}

	// Write the given arguments to the Logger object with a newline character and return the number of bytes written and any error
	return fmt.Fprintln(l.w, a...)
}

// Warningf writes the given formatted string and arguments to the Logger object if the log level is WarningLevel or higher
func (l Logger) Warningf(format string, a ...any) (int, error) {
	// If the log level is not WarningLevel or higher, return 0 and nil
	if !l.IsWarningLevel() {
		return 0, nil
	}

	// Write the given formatted string and arguments to the Logger object and return the number of bytes written and any error
	return fmt.Fprintf(l.w, format, a...)
}

// std is the global instance of the Logger struct with the default output destination and log level
var std = New(os.Stderr, WarningLevel)

// SetOutput sets the output destination for the global instance of the Logger struct
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetLevel sets the log level for the global instance of the Logger struct
func SetLevel(level Level) {
	std.SetLevel(level)
}

// IsDebugLevel returns true if the log level of the global instance of the Logger struct is DebugLevel or higher
func IsDebugLevel() bool {
	return std.IsDebugLevel()
}

// IsWarningLevel returns true if the log level of the global instance of the Logger struct is WarningLevel or higher
func IsWarningLevel() bool {
	return std.IsWarningLevel()
}

// Default returns the global instance of the Logger struct
func Default() *Logger {
	return std
}

// Debug writes the given arguments to the global instance of the Logger struct if the log level is DebugLevel or higher
func Debug(a ...any) (int, error) {
	return std.Debug(a...)
}

// Debugln writes the given arguments to the global instance of the Logger struct with a newline character if the log level is DebugLevel or higher
func Debugln(a ...any) (int, error) {
	return std.Debugln(a...)
}

// Debugf writes the given formatted string and arguments to the global instance of the Logger struct if the log level is DebugLevel or higher
func Debugf(format string, a ...any) (int, error) {
	return std.Debugf(format, a...)
}

// Warn writes the given arguments to the global instance of the Logger struct if the log level is WarningLevel or higher
func Warn(a ...any) (int, error) {
	return std.Warning(a...)
}

// Warnln writes the given arguments to the global instance of the Logger struct with a newline character if the log level is WarningLevel or higher
func Warnln(a ...any) (int, error) {
	return std.Warningln(a...)
}

// Warnf writes the given formatted string and arguments to the global instance of the Logger struct if the log level is WarningLevel or higher
func Warnf(format string, a ...any) (int, error) {
	return std.Warningf(format, a...)
}
