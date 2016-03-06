package adaptlog

// PrintLogger interface. Introduces Print logging facilities.
type PrintLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

// FatalLogger interface. Introduce Fatal logging facilities
type FatalLogger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

// PanicLogger interface. Introduce Panic logging facilities
type PanicLogger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// StdLogger interface
type StdLogger interface {
	PrintLogger
	FatalLogger
	PanicLogger
}

// StandardLogger structure
type StandardLogger struct {
	logger StdLogger
}

// Logger standard logger instance
var Logger StandardLogger

// ConfigStandardLogger configures a standard logger
func ConfigStandardLogger(logger StdLogger) {
	Logger = StandardLogger{logger}
}

// Print logging
func (l *StandardLogger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

// Printf logging with message
func (l *StandardLogger) Printf(msg string, args ...interface{}) {
	l.logger.Printf(msg, args...)
}

// Println logging with new line
func (l *StandardLogger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

// Panic logging
func (l *StandardLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Panicf logging with message
func (l *StandardLogger) Panicf(msg string, args ...interface{}) {
	l.logger.Panicf(msg, args...)
}

// Panicln logging with new line
func (l *StandardLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Fatal logging
func (l *StandardLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf logging with message
func (l *StandardLogger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}

// Fatalln logging with new line
func (l *StandardLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}
