package adaptlog

import "errors"

// ExtLogger interface. Introduces Panic and Fatal logging facilities.
type ExtLogger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

// StdLogger interface
type StdLogger interface {
	MinLogger
	ExtLogger
}

// StandardLogger structure
type StandardLogger struct {
	logger StdLogger
}

var stdLogger StdLogger

// ConfigStandardLogger configures a standard logger
func ConfigStandardLogger(logger StdLogger) {
	stdLogger = logger
}

// NewStandardLogger creates a new standard logger
func NewStandardLogger() (*StandardLogger, error) {

	if stdLogger == nil {
		return nil, errors.New("Standard logger is not configured!")
	}

	return &StandardLogger{stdLogger}, nil
}

// Print logging
func (l *StandardLogger) Print(args ...interface{}) {
	l.logger.Print(args)
}

// Printf logging with message
func (l *StandardLogger) Printf(msg string, args ...interface{}) {
	l.logger.Printf(msg, args)
}

// Println logging with new line
func (l *StandardLogger) Println(args ...interface{}) {
	l.logger.Println(args)
}

// Panic logging
func (l *StandardLogger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

// Panicf logging with message
func (l *StandardLogger) Panicf(msg string, args ...interface{}) {
	l.logger.Panicf(msg, args)
}

// Panicln logging with new line
func (l *StandardLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args)
}

// Fatal logging
func (l *StandardLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

// Fatalf logging with message
func (l *StandardLogger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args)
}

// Fatalln logging with new line
func (l *StandardLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args)
}
