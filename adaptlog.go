package adaptlog

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
