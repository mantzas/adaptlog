package adaptlog

import (
	"reflect"
	"testing"
)

func TestNewLeveledLoggerWithoutConfigReturnsError(t *testing.T) {

	logger, err := NewLevelLogger(Debug)

	if logger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewDefaultLeveledLoggerWithoutConfigReturnsError(t *testing.T) {

	logger, err := NewDefaultLevelLogger()

	if logger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewDefaultLeveledLoggerSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigLevelLogger(logger, Debug)

	stdLogger, err := NewDefaultLevelLogger()

	if stdLogger == nil {
		t.Fatal("Logger should have been not nil!")
	}

	if err != nil {
		t.Fatal("Should not have returned a error")
	}
}

func TestNewDefaultLevelLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigLevelLogger(logger, Debug)

	stdLogger, err := NewDefaultLevelLogger()

	if stdLogger == nil {
		t.Fatal("Logger should have been not nil!")
	}

	if err != nil {
		t.Fatal("Should not have returned a error")
	}

	stdLogger.Fatal()
	stdLogger.Fatalf("Test")
	stdLogger.Fatalln()

	stdLogger.Panic()
	stdLogger.Panicf("Test")
	stdLogger.Panicln()

	stdLogger.Error()
	stdLogger.Errorf("Test")
	stdLogger.Errorln()

	stdLogger.Warn()
	stdLogger.Warnf("Test")
	stdLogger.Warnln()

	stdLogger.Info()
	stdLogger.Infof("Test")
	stdLogger.Infoln()

	stdLogger.Debug()
	stdLogger.Debugf("Test")
	stdLogger.Debugln()

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
	in  Level
	out []string
}{
	{-1, []string{}},
	{Panic, []string{PanicMsg, PanicfMsg, PaniclnMsg}},
	{Fatal, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg}},
	{Error, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg}},
	{Warn, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg}},
	{Info, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg, InfoMsg, InfofMsg, InfolnMsg}},
	{Debug, []string{PanicMsg, PanicfMsg, PaniclnMsg, FatalMsg, FatalfMsg, FatallnMsg, ErrorMsg, ErrorfMsg, ErrorlnMsg, WarnMsg, WarnfMsg, WarnlnMsg, InfoMsg, InfofMsg, InfolnMsg, DebugMsg, DebugfMsg, DebuglnMsg}},
}

func TestLevelLoggingTableTest(t *testing.T) {

	for _, test := range leveledLoggingTests {

		var logger = new(TestLevelLogger)
		ConfigLevelLogger(logger, Debug)

		lvllogger, err := NewLevelLogger(test.in)

		if lvllogger == nil {
			t.Fatal("Logger should have been not nil!")
		}

		if err != nil {
			t.Fatal("Should not have returned a error")
		}

		lvllogger.Panic()
		lvllogger.Panicf("Test")
		lvllogger.Panicln()

		lvllogger.Fatal()
		lvllogger.Fatalf("Test")
		lvllogger.Fatalln()

		lvllogger.Error()
		lvllogger.Errorf("Test")
		lvllogger.Errorln()

		lvllogger.Warn()
		lvllogger.Warnf("Test")
		lvllogger.Warnln()

		lvllogger.Info()
		lvllogger.Infof("Test")
		lvllogger.Infoln()

		lvllogger.Debug()
		lvllogger.Debugf("Test")
		lvllogger.Debugln()

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
)

type TestLevelLogger struct {
	loggingData []string
}

func (l *TestLevelLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicMsg)
}

func (l *TestLevelLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, PanicfMsg)
}

func (l *TestLevelLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, PaniclnMsg)
}

func (l *TestLevelLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalMsg)
}

func (l *TestLevelLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, FatalfMsg)
}

func (l *TestLevelLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, FatallnMsg)
}

func (l *TestLevelLogger) Error(args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorMsg)
}

func (l *TestLevelLogger) Errorf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorfMsg)
}

func (l *TestLevelLogger) Errorln(args ...interface{}) {
	l.loggingData = append(l.loggingData, ErrorlnMsg)
}

func (l *TestLevelLogger) Warn(args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnMsg)
}

func (l *TestLevelLogger) Warnf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnfMsg)
}

func (l *TestLevelLogger) Warnln(args ...interface{}) {
	l.loggingData = append(l.loggingData, WarnlnMsg)
}

func (l *TestLevelLogger) Info(args ...interface{}) {
	l.loggingData = append(l.loggingData, InfoMsg)
}

func (l *TestLevelLogger) Infof(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, InfofMsg)
}

func (l *TestLevelLogger) Infoln(args ...interface{}) {
	l.loggingData = append(l.loggingData, InfolnMsg)
}

func (l *TestLevelLogger) Debug(args ...interface{}) {
	l.loggingData = append(l.loggingData, DebugMsg)
}

func (l *TestLevelLogger) Debugf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, DebugfMsg)
}

func (l *TestLevelLogger) Debugln(args ...interface{}) {
	l.loggingData = append(l.loggingData, DebuglnMsg)
}
