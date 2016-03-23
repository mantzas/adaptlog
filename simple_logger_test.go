package adaptlog

import (
	"testing"
)

func TestSimpleLoggerWithoutConfigurationIsNil(t *testing.T) {

	if Simple != nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewSimpleLoggerSucceeds(t *testing.T) {

	var logger = new(TestSimpleLogger)

	ConfigureSimpleLogger(logger)

	if Simple == nil {
		t.Fatal("Logger should have been not nil!")
	}
}

func TestNewSimpleLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestSimpleLogger)

	ConfigureSimpleLogger(logger)

	if Simple == nil {
		t.Fatal("Logger should have been not nil!")
	}

	Simple.Print("")
	Simple.Printf("Test")
	Simple.Println("")

	Simple.Fatal("")
	Simple.Fatalf("Test")
	Simple.Fatalln("")

	Simple.Panic("")
	Simple.Panicf("Test")
	Simple.Panicln("")

	if len(logger.loggingData) != 9 {
		t.Fatal("Logged items should be 9!")
	}

	if logger.loggingData[0] != PrintMsg || logger.loggingData[1] != PrintfMsg || logger.loggingData[2] != PrintlnMsg ||
		logger.loggingData[3] != FatalMsg || logger.loggingData[4] != FatalfMsg || logger.loggingData[5] != FatallnMsg ||
		logger.loggingData[6] != PanicMsg || logger.loggingData[7] != PanicfMsg || logger.loggingData[8] != PaniclnMsg {
		t.Fatal("Logged items do not match!")
	}
}

const (
	PanicMsg   = "PanicMsg"
	PanicfMsg  = "PanicfMsg"
	PaniclnMsg = "PaniclnMsg"

	FatalMsg   = "FatalMsg"
	FatalfMsg  = "FatalfMsg"
	FatallnMsg = "FatallnMsg"

	PrintMsg   = "PrintMsg"
	PrintfMsg  = "PrintfMsg"
	PrintlnMsg = "PrintlnMsg"
)

type TestSimpleLogger struct {
	loggingData []string
}

func (l *TestSimpleLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicMsg)
}

func (l *TestSimpleLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicfMsg)
}

func (l *TestSimpleLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, PaniclnMsg)
}

func (l *TestSimpleLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalMsg)
}

func (l *TestSimpleLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalfMsg)
}

func (l *TestSimpleLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatallnMsg)
}

func (l *TestSimpleLogger) Print(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintMsg)
}

func (l *TestSimpleLogger) Printf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintfMsg)
}

func (l *TestSimpleLogger) Println(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintlnMsg)
}
