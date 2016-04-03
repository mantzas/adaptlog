package adaptlog

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// LogLevel enum
type LogLevel uint8

// Log Level values
const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarnLevel:
		return "Warn"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	default:
		return fmt.Sprintf("Not mapped value %d", l)
	}
}

var logFunc func(LogLevel, string, ...interface{})
var loglnFunc func(LogLevel, string, ...interface{})
var logfFunc func(LogLevel, string, string, ...interface{})

// StdLevelLogger is a level logger based on the standard log package
type StdLevelLogger struct {
	logFunc   func(LogLevel, string, ...interface{})
	loglnFunc func(LogLevel, string, ...interface{})
	logfFunc  func(LogLevel, string, string, ...interface{})
	context   string
}

// NewStdLevelLogger creates a new standard level logger
func NewStdLevelLogger(context string) *StdLevelLogger {
	return &StdLevelLogger{logFunc, loglnFunc, logfFunc, context}
}

// ConfigureStdLevelLogger configures the standard level logger with the default log level.
// By providing a different io.Writer and prefix you can control the logging output (testing etc)
func ConfigureStdLevelLogger(defaultLoglevel LogLevel, w io.Writer, context string) {

	logFunc = func(level LogLevel, context string, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}

		if context == "" {
			args = prepend(args, strings.Join([]string{level.String(), " "}, ""))
		} else {
			args = prepend(args, level.String(), " ", context, " ")
		}

		log.Print(args...)
	}

	loglnFunc = func(level LogLevel, context string, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}

		if context == "" {
			args = prepend(args, level.String())
		} else {
			args = prepend(args, level.String(), context)
		}

		log.Println(args...)
	}

	logfFunc = func(level LogLevel, context string, msg string, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}

		if context == "" {
			log.Printf(strings.Join([]string{level.String(), msg}, " "), args...)
		} else {
			log.Printf(strings.Join([]string{level.String(), context, msg}, " "), args...)
		}
	}

	Level = NewStdLevelLogger(context)
}

func prepend(args []interface{}, values ...string) []interface{} {
	valuesLength := len(values)
	argsLength := len(args)
	argsExt := make([]interface{}, argsLength+valuesLength)

	for index := 0; index < valuesLength; index++ {
		argsExt[index] = values[index]
	}

	for index := valuesLength; index < argsLength+valuesLength; index++ {
		argsExt[index] = args[index-valuesLength]
	}
	return argsExt
}

// Fatal logging
func (l *StdLevelLogger) Fatal(args ...interface{}) {

	l.logFunc(FatalLevel, l.context, args...)
}

// Fatalf logging
func (l *StdLevelLogger) Fatalf(msg string, args ...interface{}) {

	l.logfFunc(FatalLevel, l.context, msg, args...)
}

// Fatalln logging
func (l *StdLevelLogger) Fatalln(args ...interface{}) {

	l.loglnFunc(FatalLevel, l.context, args...)
}

// Error logging
func (l *StdLevelLogger) Error(args ...interface{}) {

	l.logFunc(ErrorLevel, l.context, args...)
}

// Errorf logging
func (l *StdLevelLogger) Errorf(msg string, args ...interface{}) {

	l.logfFunc(ErrorLevel, l.context, msg, args...)
}

// Errorln logging
func (l *StdLevelLogger) Errorln(args ...interface{}) {

	l.loglnFunc(ErrorLevel, l.context, args...)
}

// Warn logging
func (l *StdLevelLogger) Warn(args ...interface{}) {

	l.logFunc(WarnLevel, l.context, args...)
}

// Warnf logging
func (l *StdLevelLogger) Warnf(msg string, args ...interface{}) {

	l.logfFunc(WarnLevel, l.context, msg, args...)
}

// Warnln logging
func (l *StdLevelLogger) Warnln(args ...interface{}) {

	l.loglnFunc(WarnLevel, l.context, args...)
}

// Info logging
func (l *StdLevelLogger) Info(args ...interface{}) {

	l.logFunc(InfoLevel, l.context, args...)
}

// Infof logging
func (l *StdLevelLogger) Infof(msg string, args ...interface{}) {

	l.logfFunc(InfoLevel, l.context, msg, args...)
}

// Infoln logging
func (l *StdLevelLogger) Infoln(args ...interface{}) {

	l.loglnFunc(InfoLevel, l.context, args...)
}

// Debug logging
func (l *StdLevelLogger) Debug(args ...interface{}) {

	l.logFunc(DebugLevel, l.context, args...)
}

// Debugf logging
func (l *StdLevelLogger) Debugf(msg string, args ...interface{}) {

	l.logfFunc(DebugLevel, l.context, msg, args...)
}

// Debugln logging
func (l *StdLevelLogger) Debugln(args ...interface{}) {

	l.loglnFunc(DebugLevel, l.context, args...)
}
