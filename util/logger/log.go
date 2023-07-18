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

func New(w io.Writer, l Level) *Logger {
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

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) IsDebugLevel() bool {
	return l.level >= DebugLevel
}

func (l *Logger) IsWarningLevel() bool {
	return l.level >= WarningLevel
}

func (l Logger) Debug(a ...any) (int, error) {
	if !l.IsDebugLevel() {
		return 0, nil
	}
	return fmt.Fprint(l.w, a...)
}

func (l Logger) Debugln(a ...any) (int, error) {
	if !l.IsDebugLevel() {
		return 0, nil
	}
	return fmt.Fprintln(l.w, a...)
}

func (l Logger) Debugf(format string, a ...any) (int, error) {
	if !l.IsDebugLevel() {
		return 0, nil
	}
	return fmt.Fprintf(l.w, format, a...)
}

func (l Logger) Warning(a ...any) (int, error) {
	if !l.IsWarningLevel() {
		return 0, nil
	}
	return fmt.Fprint(l.w, a...)
}

func (l Logger) Warningln(a ...any) (int, error) {
	if !l.IsWarningLevel() {
		return 0, nil
	}
	return fmt.Fprintln(l.w, a...)
}

func (l Logger) Warningf(format string, a ...any) (int, error) {
	if !l.IsWarningLevel() {
		return 0, nil
	}
	return fmt.Fprintf(l.w, format, a...)
}

var std = New(os.Stderr, WarningLevel)

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

func IsDebugLevel() bool {
	return std.IsDebugLevel()
}

func IsWarningLevel() bool {
	return std.IsWarningLevel()
}

func Default() *Logger {
	return std
}

func Debug(a ...any) (int, error) {
	return std.Debug(a...)
}

func Debugln(a ...any) (int, error) {
	return std.Debugln(a...)
}

func Debugf(format string, a ...any) (int, error) {
	return std.Debugf(format, a...)
}

func Warn(a ...any) (int, error) {
	return std.Warning(a...)
}

func Warnln(a ...any) (int, error) {
	return std.Warningln(a...)
}

func Warnf(format string, a ...any) (int, error) {
	return std.Warningf(format, a...)
}
