package adaptlog

// LogLevel type
type LogLevel uint8

// Logging levels.
const (
	// Panic
	Panic LogLevel = iota
	// Fatal
	Fatal
	// Error
	Error
	// Warn
	Warn
	// Info
	Info
	// Debug
	Debug
)

// LevelLogger interface
type LevelLogger interface {
	ExtendedLogger
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
}
