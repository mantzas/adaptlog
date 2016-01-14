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

// StdLogger interface
type StdLogger interface {
	BaseLogger
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// LevelLogger interface
type LevelLogger interface {
	StdLogger
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
