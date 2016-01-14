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

// DefaultLogLevel defines the global log level if not specified explicitely
var DefaultLogLevel LogLevel

type levelLogger struct {
	Level  LogLevel
	logger LevelLogger
}

func (l *levelLogger) Panic(args ...interface{}) {
	if l.Level > Panic {
		return
	}

	l.logger.Panic(args)
}

func (l *levelLogger) Panicf(msg string, args ...interface{}) {
	if l.Level < Panic {
		return
	}
	l.logger.Panicf(msg, args)
}

func (l *levelLogger) Panicln(args ...interface{}) {
	if l.Level < Panic {
		return
	}
	l.logger.Panicln(args)
}

func (l *levelLogger) Fatal(args ...interface{}) {
	if l.Level < Fatal {
		return
	}
	l.logger.Fatal(args)
}

func (l *levelLogger) Fatalf(msg string, args ...interface{}) {
	if l.Level < Fatal {
		return
	}
	l.logger.Fatalf(msg, args)
}

func (l *levelLogger) Fatalln(args ...interface{}) {
	if l.Level < Fatal {
		return
	}
	l.logger.Fatalln(args)
}

func (l *levelLogger) Error(args ...interface{}) {
	if l.Level < Error {
		return
	}
	l.logger.Error(args)
}

func (l *levelLogger) Errorf(msg string, args ...interface{}) {
	if l.Level < Error {
		return
	}
	l.logger.Errorf(msg, args)
}

func (l *levelLogger) Errorln(args ...interface{}) {
	if l.Level < Error {
		return
	}
	l.logger.Errorln(args)
}

func (l *levelLogger) Warn(args ...interface{}) {
	if l.Level < Warn {
		return
	}
	l.logger.Warn(args)
}

func (l *levelLogger) Warnf(msg string, args ...interface{}) {
	if l.Level < Warn {
		return
	}
	l.logger.Warnf(msg, args)
}

func (l *levelLogger) Warnln(args ...interface{}) {
	if l.Level < Warn {
		return
	}
	l.logger.Warnln(args)
}

func (l *levelLogger) Info(args ...interface{}) {
	if l.Level < Info {
		return
	}
	l.logger.Info(args)
}

func (l *levelLogger) Infof(msg string, args ...interface{}) {
	if l.Level < Info {
		return
	}
	l.logger.Infof(msg, args)
}

func (l *levelLogger) Infoln(args ...interface{}) {
	if l.Level < Info {
		return
	}
	l.logger.Infoln(args)
}

func (l *levelLogger) Debug(args ...interface{}) {
	if l.Level < Debug {
		return
	}
	l.logger.Debug(args)
}

func (l *levelLogger) Debugf(msg string, args ...interface{}) {
	if l.Level < Debug {
		return
	}
	l.logger.Debugf(msg, args)
}

func (l *levelLogger) Debugln(args ...interface{}) {
	if l.Level < Debug {
		return
	}
	l.logger.Debugln(args)
}
