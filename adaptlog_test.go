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

	stdLogger.Print("")
	stdLogger.Printf("Test")
	stdLogger.Println("")

	stdLogger.Fatal("")
	stdLogger.Fatalf("Test")
	stdLogger.Fatalln("")

	stdLogger.Panic("")
	stdLogger.Panicf("Test")
	stdLogger.Panicln("")

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
	PrintMsg   = "PrintMsg"
	PrintfMsg  = "PrintfMsg"
	PrintlnMsg = "PrintlnMsg"

	PanicMsg   = "PanicMsg"
	PanicfMsg  = "PanicfMsg"
	PaniclnMsg = "PaniclnMsg"

	FatalMsg   = "FatalMsg"
	FatalfMsg  = "FatalfMsg"
	FatallnMsg = "FatallnMsg"
)

type TestStandardLogger struct {
	loggingData []string
}

func (l *TestStandardLogger) Print(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintMsg)
}

func (l *TestStandardLogger) Printf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintfMsg)
}

func (l *TestStandardLogger) Println(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintlnMsg)
}

func (l *TestStandardLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicMsg)
}

func (l *TestStandardLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicfMsg)
}

func (l *TestStandardLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, PaniclnMsg)
}

func (l *TestStandardLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalMsg)
}

func (l *TestStandardLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalfMsg)
}

func (l *TestStandardLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatallnMsg)
}
