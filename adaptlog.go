package adaptlog

import "errors"

// MinLogger interface. Introduces Print logging facilities.
type MinLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

// MinimalLogger structure
type MinimalLogger struct {
	logger MinLogger
}

var minLogger MinLogger

// ConfigMinimalLogger configures a minimal logger
func ConfigMinimalLogger(logger MinLogger) {
	minLogger = logger
}

// NewMinimumLogger creates a new minimal logger
func NewMinimumLogger() (*MinimalLogger, error) {

	if minLogger == nil {
		return nil, errors.New("Minimal logger is not configured!")
	}

	return &MinimalLogger{minLogger}, nil
}

// Print logging
func (l *MinimalLogger) Print(args ...interface{}) {
	l.logger.Print(args)
}

// Printf logging with message
func (l *MinimalLogger) Printf(msg string, args ...interface{}) {
	l.logger.Printf(msg, args)
}

// Println logging with new line
func (l *MinimalLogger) Println(args ...interface{}) {
	l.logger.Println(args)
}
