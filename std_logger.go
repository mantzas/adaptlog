package adaptlog

// ExtLogger interface. Introduces Panic and Fatal logging facilities.
type ExtLogger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

// StdLogger interface
type StdLogger interface {
	MinLogger
	ExtLogger
}

// StandardLogger structure
type StandardLogger struct {
}
