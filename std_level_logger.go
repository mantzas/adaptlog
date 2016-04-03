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

var logFunc func(LogLevel, ...interface{})
var loglnFunc func(LogLevel, ...interface{})
var logfFunc func(LogLevel, string, ...interface{})

// ConfigureStdLevelLogger configures the standard level logger with the default log level.
// By providing a different io.Writer and prefix you can control the logging output (testing etc)
func ConfigureStdLevelLogger(defaultLoglevel LogLevel, w io.Writer) {

	//logLevel := strings.Join([]string{level.String(), " "}, "")
	logFunc = func(level LogLevel, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}
		args = prepend(args, level.String())
		log.Print(args...)
	}

	loglnFunc = func(level LogLevel, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}
		args = prepend(args, level.String())
		log.Println(args...)
	}

	logfFunc = func(level LogLevel, msg string, args ...interface{}) {
		if defaultLoglevel > level {
			return
		}

		if w != nil {
			log.SetOutput(w)
		}

		log.Printf(strings.Join([]string{level.String(), msg}, " "), args...)
	}
}

func prepend(args []interface{}, value string) []interface{} {
	argsExt := make([]interface{}, len(args)+1)
	argsExt[0] = value
	for index := 0; index < len(args); index++ {
		argsExt[index+1] = args[index]
	}
	return argsExt
}

// StdLevelLogger is a level logger based on the standard log package
type StdLevelLogger struct {
	logFunc   func(LogLevel, ...interface{})
	loglnFunc func(LogLevel, ...interface{})
	logfFunc  func(LogLevel, string, ...interface{})
}

// NewStdLevelLogger creates a new standard level logger
func NewStdLevelLogger() LevelLogger {

	return &StdLevelLogger{logFunc, loglnFunc, logfFunc}
}

// Fatal logging
func (l *StdLevelLogger) Fatal(args ...interface{}) {

	l.logFunc(FatalLevel, args...)
}

// Fatalf logging
func (l *StdLevelLogger) Fatalf(msg string, args ...interface{}) {

	l.logfFunc(FatalLevel, msg, args...)
}

// Fatalln logging
func (l *StdLevelLogger) Fatalln(args ...interface{}) {

	l.loglnFunc(FatalLevel, args...)
}

// Error logging
func (l *StdLevelLogger) Error(args ...interface{}) {

	l.logFunc(ErrorLevel, args...)
}

// Errorf logging
func (l *StdLevelLogger) Errorf(msg string, args ...interface{}) {

	l.logfFunc(ErrorLevel, msg, args...)
}

// Errorln logging
func (l *StdLevelLogger) Errorln(args ...interface{}) {

	l.loglnFunc(ErrorLevel, args...)
}

// Warn logging
func (l *StdLevelLogger) Warn(args ...interface{}) {

	l.logFunc(WarnLevel, args...)
}

// Warnf logging
func (l *StdLevelLogger) Warnf(msg string, args ...interface{}) {

	l.logfFunc(WarnLevel, msg, args...)
}

// Warnln logging
func (l *StdLevelLogger) Warnln(args ...interface{}) {

	l.loglnFunc(WarnLevel, args...)
}

// Info logging
func (l *StdLevelLogger) Info(args ...interface{}) {

	l.logFunc(InfoLevel, args...)
}

// Infof logging
func (l *StdLevelLogger) Infof(msg string, args ...interface{}) {

	l.logfFunc(InfoLevel, msg, args...)
}

// Infoln logging
func (l *StdLevelLogger) Infoln(args ...interface{}) {

	l.loglnFunc(InfoLevel, args...)
}

// Debug logging
func (l *StdLevelLogger) Debug(args ...interface{}) {

	l.logFunc(DebugLevel, args...)
}

// Debugf logging
func (l *StdLevelLogger) Debugf(msg string, args ...interface{}) {

	l.logfFunc(DebugLevel, msg, args...)
}

// Debugln logging
func (l *StdLevelLogger) Debugln(args ...interface{}) {

	l.loglnFunc(DebugLevel, args...)
}
