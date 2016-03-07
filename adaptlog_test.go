package adaptlog

import (
	"reflect"
	"testing"
)

func TestLoggerWithoutConfigurationIsNil(t *testing.T) {

	if log != nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewLoggerSucceeds(t *testing.T) {

	var logger = new(TestLogger)

	Configure(logger, AnyLevel)

	if log == nil {
		t.Fatal("Logger should have been not nil!")
	}
}

func TestNewStandardLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLogger)

	Configure(logger, AnyLevel)

	if log == nil {
		t.Fatal("Logger should have been not nil!")
	}

	Print("")
	Printf("Test")
	Println("")

	Fatal("")
	Fatalf("Test")
	Fatalln("")

	Panic("")
	Panicf("Test")
	Panicln("")

	if len(logger.loggingData) != 9 {
		t.Fatal("Logged items should be 9!")
	}

	if logger.loggingData[0] != PrintMsg || logger.loggingData[1] != PrintfMsg || logger.loggingData[2] != PrintlnMsg ||
		logger.loggingData[3] != FatalMsg || logger.loggingData[4] != FatalfMsg || logger.loggingData[5] != FatallnMsg ||
		logger.loggingData[6] != PanicMsg || logger.loggingData[7] != PanicfMsg || logger.loggingData[8] != PaniclnMsg {
		t.Fatal("Logged items do not match!")
	}
}

func TestLeveledLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLogger)

	Configure(logger, DebugLevel)

	if log == nil {
		t.Fatal("Logger should have been not nil!")
	}

	Fatal("")
	Fatalf("Test")
	Fatalln("")

	Panic("")
	Panicf("Test")
	Panicln("")

	Error("")
	Errorf("Test")
	Errorln("")

	Warn("")
	Warnf("Test")
	Warnln("")

	Info("")
	Infof("Test")
	Infoln("")

	Debug("")
	Debugf("Test")
	Debugln("")

	if len(logger.loggingData) != 18 {
		t.Fatal("Logged items should be 18!")
	}

	if logger.loggingData[0] != FatalMsg || logger.loggingData[1] != FatalfMsg || logger.loggingData[2] != FatallnMsg ||
		logger.loggingData[3] != PanicMsg || logger.loggingData[4] != PanicfMsg || logger.loggingData[5] != PaniclnMsg ||
		logger.loggingData[6] != ErrorMsg || logger.loggingData[7] != ErrorfMsg || logger.loggingData[8] != ErrorlnMsg ||
		logger.loggingData[9] != WarnMsg || logger.loggingData[10] != WarnfMsg || logger.loggingData[11] != WarnlnMsg ||
		logger.loggingData[12] != InfoMsg || logger.loggingData[13] != InfofMsg || logger.loggingData[14] != InfolnMsg ||
		logger.loggingData[15] != DebugMsg || logger.loggingData[16] != DebugfMsg || logger.loggingData[17] != DebuglnMsg {
		t.Fatal("Logged items do not match!")
	}
}

var leveledLoggingTests = []struct {
	in  LogLevel
	out []string
}{
	{-1, []string{}},
	{PanicLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg}},
	{FatalLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg}},
	{ErrorLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg}},
	{WarnLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg}},
	{InfoLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg, InfoMsg, InfofMsg, InfolnMsg}},
	{DebugLevel, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg, InfoMsg, InfofMsg, InfolnMsg, DebugMsg, DebugfMsg, DebuglnMsg}},
}

func TestLevelLoggingTableTest(t *testing.T) {

	for _, test := range leveledLoggingTests {

		var logger = new(TestLogger)
		Configure(logger, test.in)

		if log == nil {
			t.Fatal("Logger should have been not nil!")
		}

		Panic("")
		Panicf("Test")
		Panicln("")

		Fatal("")
		Fatalf("Test")
		Fatalln("")

		Error("")
		Errorf("Test")
		Errorln("")

		Warn("")
		Warnf("Test")
		Warnln("")

		Info("")
		Infof("Test")
		Infoln("")

		Debug("")
		Debugf("Test")
		Debugln("")

		if len(logger.loggingData) == len(test.out) && len(test.out) == 0 {
			if test.in != -1 {
				t.Fatal("Level should have been -1!")
			}
		} else {
			if !reflect.DeepEqual(logger.loggingData, test.out) {
				t.Fatalf("Logged data is different as expected! actual=%s expected=%s", logger.loggingData, test.out)
			}
		}
	}
}

const (
	PanicMsg   = "PanicMsg"
	PanicfMsg  = "PanicfMsg"
	PaniclnMsg = "PaniclnMsg"

	FatalMsg   = "FatalMsg"
	FatalfMsg  = "FatalfMsg"
	FatallnMsg = "FatallnMsg"

	ErrorMsg   = "ErrorMsg"
	ErrorfMsg  = "ErrorfMsg"
	ErrorlnMsg = "ErrorlnMsg"

	WarnMsg   = "WarnMsg"
	WarnfMsg  = "WarnfMsg"
	WarnlnMsg = "WarnlnMsg"

	InfoMsg   = "InfoMsg"
	InfofMsg  = "InfofMsg"
	InfolnMsg = "InfolnMsg"

	DebugMsg   = "DebugMsg"
	DebugfMsg  = "DebugfMsg"
	DebuglnMsg = "DebuglnMsg"

	PrintMsg   = "PrintMsg"
	PrintfMsg  = "PrintfMsg"
	PrintlnMsg = "PrintlnMsg"
)

type TestLogger struct {
	loggingData []string
}

func (l *TestLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicMsg)
}

func (l *TestLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicfMsg)
}

func (l *TestLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, PaniclnMsg)
}

func (l *TestLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalMsg)
}

func (l *TestLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalfMsg)
}

func (l *TestLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatallnMsg)
}

func (l *TestLogger) Error(args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorMsg)
}

func (l *TestLogger) Errorf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorfMsg)
}

func (l *TestLogger) Errorln(args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorlnMsg)
}

func (l *TestLogger) Warn(args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnMsg)
}

func (l *TestLogger) Warnf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnfMsg)
}

func (l *TestLogger) Warnln(args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnlnMsg)
}

func (l *TestLogger) Info(args ...interface{}) {
	l.loggingData = append(l.loggingData, InfoMsg)
}

func (l *TestLogger) Infof(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, InfofMsg)
}

func (l *TestLogger) Infoln(args ...interface{}) {
	l.loggingData = append(l.loggingData, InfolnMsg)
}

func (l *TestLogger) Debug(args ...interface{}) {
	l.loggingData = append(l.loggingData, DebugMsg)
}

func (l *TestLogger) Debugf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, DebugfMsg)
}

func (l *TestLogger) Debugln(args ...interface{}) {
	l.loggingData = append(l.loggingData, DebuglnMsg)
}

func (l *TestLogger) Print(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintMsg)
}

func (l *TestLogger) Printf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintfMsg)
}

func (l *TestLogger) Println(args ...interface{}) {
	l.loggingData = append(l.loggingData, PrintlnMsg)
}
