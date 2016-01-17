package adaptlog

import "testing"

func TestNewStandardLoggerWithoutConfigReturnsError(t *testing.T) {

	logger, err := NewStandardLogger()

	if logger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewStandardLoggerSucceeds(t *testing.T) {

	var logger = new(TestStandardLogger)

	ConfigStandardLogger(logger)

	stdLogger, err := NewStandardLogger()

	if stdLogger == nil {
		t.Fatal("Logger should have been not nil!")
	}

	if err != nil {
		t.Fatal("Should not have returned a error")
	}
}

func TestNewStandardLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestStandardLogger)

	ConfigStandardLogger(logger)

	stdLogger, err := NewStandardLogger()

	if stdLogger == nil {
		t.Fatal("Logger should have been not nil!")
	}

	if err != nil {
		t.Fatal("Should not have returned a error")
	}

	stdLogger.Print()
	stdLogger.Printf("Test")
	stdLogger.Println()

	stdLogger.Fatal()
	stdLogger.Fatalf("Test")
	stdLogger.Fatalln()

	stdLogger.Panic()
	stdLogger.Panicf("Test")
	stdLogger.Panicln()

	if len(logger.loggingData) != 9 {
		t.Fatal("Logged items should be 9!")
	}

	if logger.loggingData[0] != Print || logger.loggingData[1] != Printf || logger.loggingData[2] != Println ||
		logger.loggingData[3] != Fatal || logger.loggingData[4] != Fatalf || logger.loggingData[5] != Fatalln ||
		logger.loggingData[6] != Panic || logger.loggingData[7] != Panicf || logger.loggingData[8] != Panicln {
		t.Fatal("Logged items do not match!")
	}
}

const (
	Panic   = "Panic"
	Panicf  = "Panicf"
	Panicln = "Panicln"

	Fatal   = "Fatal"
	Fatalf  = "Fatalf"
	Fatalln = "Fatalln"
)

type TestStandardLogger struct {
	loggingData []string
}

func (l *TestStandardLogger) Print(args ...interface{}) {
	l.loggingData = append(l.loggingData, Print)
}

func (l *TestStandardLogger) Printf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Printf)
}

func (l *TestStandardLogger) Println(args ...interface{}) {
	l.loggingData = append(l.loggingData, Println)
}

func (l *TestStandardLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, Panic)
}

func (l *TestStandardLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Panicf)
}

func (l *TestStandardLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Panicln)
}

func (l *TestStandardLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatal)
}

func (l *TestStandardLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatalf)
}

func (l *TestStandardLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatalln)
}
