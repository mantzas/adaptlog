package adaptlog

// SimpleLogger the simple logger interface with Print, Panic, Fatal
type SimpleLogger interface {
	PrintLogger
	FatalLogger
	PanicLogger
}

// Simple Logger
var Simple SimpleLogger

// ConfigureSimpleLogger configures the simple logger
func ConfigureSimpleLogger(logger SimpleLogger) {
	Simple = logger
}
