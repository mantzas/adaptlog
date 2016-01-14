package adaptlog

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
