package adaptlog

import (
	"testing"
)

func TestLeveledLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigureLevelLogger(logger)

	if Level == nil {
		t.Fatal("Logger should have been not nil!")
	}

	Level.Fatal("")
	Level.Fatalf("Test")
	Level.Fatalln("")

	Level.Error("")
	Level.Errorf("Test")
	Level.Errorln("")

	Level.Warn("")
	Level.Warnf("Test")
	Level.Warnln("")

	Level.Info("")
	Level.Infof("Test")
	Level.Infoln("")

	Level.Debug("")
	Level.Debugf("Test")
	Level.Debugln("")

	if len(logger.loggingData) != 15 {
		t.Fatal("Logged items should be 15!")
	}

	if logger.loggingData[0] != FatalMsg || logger.loggingData[1] != FatalfMsg || logger.loggingData[2] != FatallnMsg ||
		logger.loggingData[3] != ErrorMsg || logger.loggingData[4] != ErrorfMsg || logger.loggingData[5] != ErrorlnMsg ||
		logger.loggingData[6] != WarnMsg || logger.loggingData[7] != WarnfMsg || logger.loggingData[8] != WarnlnMsg ||
		logger.loggingData[9] != InfoMsg || logger.loggingData[10] != InfofMsg || logger.loggingData[11] != InfolnMsg ||
		logger.loggingData[12] != DebugMsg || logger.loggingData[13] != DebugfMsg || logger.loggingData[14] != DebuglnMsg {
		t.Fatal("Logged items do not match!")
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
