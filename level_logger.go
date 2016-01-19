package adaptlog

import "errors"

// ErrorLogger interface. Introduces Error logging facilities.
type ErrorLogger interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
}

// WarnLogger interface. Introduces Warn logging facilities.
type WarnLogger interface {
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})
}

// InfoLogger interface. Introduces Info logging facilities.
type InfoLogger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
}

// DebugLogger interface. Introduces Debug logging facilities.
type DebugLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
}

// LvlLogger interface. Introduces Error, Warn, Info and Debug logging facilities.
type LvlLogger interface {
	FatalLogger
	PanicLogger
	ErrorLogger
	WarnLogger
	InfoLogger
	DebugLogger
}

var levelLogger LvlLogger

// ConfigLevelLogger configures a leveled logger
func ConfigLevelLogger(logger LvlLogger) {
	levelLogger = logger
}

// NewLevelLogger creates a new level logger
func NewLevelLogger() (*LevelLogger, error) {

	if levelLogger == nil {
		return nil, errors.New("Level logger is not configured!")
	}

	return &LevelLogger{levelLogger}, nil
}

// LevelLogger for logging with level support
type LevelLogger struct {
	logger LvlLogger
}

// Panic level logging
func (l *LevelLogger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

// Panicf level logging with message
func (l *LevelLogger) Panicf(msg string, args ...interface{}) {
	l.logger.Panicf(msg, args)
}

// Panicln level logging with new line
func (l *LevelLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args)
}

// Fatal level logging
func (l *LevelLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

// Fatalf level logging with message
func (l *LevelLogger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args)
}

// Fatalln level logging with new line
func (l *LevelLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args)
}

// Error level logging
func (l *LevelLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}

// Errorf level logging with message
func (l *LevelLogger) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args)
}

// Errorln level logging with new line
func (l *LevelLogger) Errorln(args ...interface{}) {
	l.logger.Errorln(args)
}

// Warn level logging
func (l *LevelLogger) Warn(args ...interface{}) {
	l.logger.Warn(args)
}

// Warnf level logging with message
func (l *LevelLogger) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args)
}

// Warnln level logging with new line
func (l *LevelLogger) Warnln(args ...interface{}) {
	l.logger.Warnln(args)
}

// Info level logging
func (l *LevelLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

// Infof level logging with message
func (l *LevelLogger) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args)
}

// Infoln level logging with new line
func (l *LevelLogger) Infoln(args ...interface{}) {
	l.logger.Infoln(args)
}

// Debug level logging
func (l *LevelLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

// Debugf level logging with message
func (l *LevelLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args)
}

// Debugln level logging with new line
func (l *LevelLogger) Debugln(args ...interface{}) {
	l.logger.Debugln(args)
}
