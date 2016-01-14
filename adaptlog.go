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

// BaseLogger interface
type BaseLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

// ExtendedLogger interface
type ExtendedLogger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

// StdLogger interface
type StdLogger interface {
	BaseLogger
	ExtendedLogger
}

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
