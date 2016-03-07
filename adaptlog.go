package adaptlog

var level LogLevel
var log Logger

// Configure the internal logger
func Configure(logger Logger, logLevel LogLevel) {
	log = logger
	level = logLevel
}

// LogLevel type
type LogLevel int

// These logging levels are used to setup the logging
const (
	// Panic level
	PanicLevel LogLevel = iota
	// Fatal level
	FatalLevel
	// Error level
	ErrorLevel
	// Warn level
	WarnLevel
	// Info level
	InfoLevel
	// Debug level
	DebugLevel
	// Any level
	AnyLevel
)

// Logger interface. Introduces Error, Warn, Info and Debug logging facilities.
type Logger interface {
	FatalLogger
	PanicLogger
	ErrorLogger
	WarnLogger
	InfoLogger
	DebugLogger
	PrintLogger
}

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

// Print logging
func Print(args ...interface{}) {
	log.Print(args...)
}

// Printf logging with message
func Printf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Println logging with new line
func Println(args ...interface{}) {
	log.Println(args...)
}

// Panic level logging
func Panic(args ...interface{}) {
	if level < PanicLevel {
		return
	}
	log.Panic(args...)
}

// Panicf level logging with message
func Panicf(msg string, args ...interface{}) {
	if level < PanicLevel {
		return
	}
	log.Panicf(msg, args...)
}

// Panicln level logging with new line
func Panicln(args ...interface{}) {
	if level < PanicLevel {
		return
	}
	log.Panicln(args...)
}

// Fatal level logging
func Fatal(args ...interface{}) {
	if level < FatalLevel {
		return
	}
	log.Fatal(args...)
}

// Fatalf level logging with message
func Fatalf(msg string, args ...interface{}) {
	if level < FatalLevel {
		return
	}
	log.Fatalf(msg, args...)
}

// Fatalln level logging with new line
func Fatalln(args ...interface{}) {
	if level < FatalLevel {
		return
	}
	log.Fatalln(args...)
}

// Error level logging
func Error(args ...interface{}) {
	if level < ErrorLevel {
		return
	}
	log.Error(args...)
}

// Errorf level logging with message
func Errorf(msg string, args ...interface{}) {
	if level < ErrorLevel {
		return
	}
	log.Errorf(msg, args...)
}

// Errorln level logging with new line
func Errorln(args ...interface{}) {
	if level < ErrorLevel {
		return
	}
	log.Errorln(args...)
}

// Warn level logging
func Warn(args ...interface{}) {
	if level < WarnLevel {
		return
	}
	log.Warn(args...)
}

// Warnf level logging with message
func Warnf(msg string, args ...interface{}) {
	if level < WarnLevel {
		return
	}
	log.Warnf(msg, args...)
}

// Warnln level logging with new line
func Warnln(args ...interface{}) {
	if level < WarnLevel {
		return
	}
	log.Warnln(args...)
}

// Info level logging
func Info(args ...interface{}) {
	if level < InfoLevel {
		return
	}
	log.Info(args...)
}

// Infof level logging with message
func Infof(msg string, args ...interface{}) {
	if level < InfoLevel {
		return
	}
	log.Infof(msg, args...)
}

// Infoln level logging with new line
func Infoln(args ...interface{}) {
	if level < InfoLevel {
		return
	}
	log.Infoln(args...)
}

// Debug level logging
func Debug(args ...interface{}) {
	if level < DebugLevel {
		return
	}
	log.Debug(args...)
}

// Debugf level logging with message
func Debugf(msg string, args ...interface{}) {
	if level < DebugLevel {
		return
	}
	log.Debugf(msg, args...)
}

// Debugln level logging with new line
func Debugln(args ...interface{}) {
	if level < DebugLevel {
		return
	}
	log.Debugln(args...)
}
