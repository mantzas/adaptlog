package adaptlog

// LevelLogger is the level logger interface
type LevelLogger interface {
	FatalLogger
	ErrorLogger
	WarnLogger
	InfoLogger
	DebugLogger
}

// Level Logger
var Level LevelLogger

// ConfigureLevelLogger configures the level logger
func ConfigureLevelLogger(logger LevelLogger) {
	Level = logger
}
