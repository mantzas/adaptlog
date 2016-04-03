package adaptlog

import (
	"fmt"
	"strings"
	"testing"
)

func TestLeveledLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigureLevelLogger(logger)

	if Level == nil {
		t.Fatal("Logger should have been not nil!")
	}

	Level.Fatal(FatalMsg)
	Level.Fatalf("Test %s", FatalfMsg)
	Level.Fatalln(FatallnMsg)

	Level.Error(ErrorMsg)
	Level.Fatalf("Test %s", ErrorfMsg)
	Level.Errorln(ErrorlnMsg)

	Level.Warn(WarnMsg)
	Level.Fatalf("Test %s", WarnfMsg)
	Level.Warnln(WarnlnMsg)

	Level.Info(InfoMsg)
	Level.Fatalf("Test %s", InfofMsg)
	Level.Infoln(InfolnMsg)

	Level.Debug(DebugMsg)
	Level.Fatalf("Test %s", DebugfMsg)
	Level.Debugln(DebuglnMsg)

	if len(logger.loggingData) != 15 {
		t.Fatal("Logged items should be 15!")
	}

	if logger.loggingData[0] != FatalMsg || logger.loggingData[1] != "Test "+FatalfMsg || logger.loggingData[2] != FatallnMsg ||
		logger.loggingData[3] != ErrorMsg || logger.loggingData[4] != "Test "+ErrorfMsg || logger.loggingData[5] != ErrorlnMsg ||
		logger.loggingData[6] != WarnMsg || logger.loggingData[7] != "Test "+WarnfMsg || logger.loggingData[8] != WarnlnMsg ||
		logger.loggingData[9] != InfoMsg || logger.loggingData[10] != "Test "+InfofMsg || logger.loggingData[11] != InfolnMsg ||
		logger.loggingData[12] != DebugMsg || logger.loggingData[13] != "Test "+DebugfMsg || logger.loggingData[14] != DebuglnMsg {
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
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoinWithMessage(msg, args...))
}

func (l *TestLevelLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Error(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Errorf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoinWithMessage(msg, args...))
}

func (l *TestLevelLogger) Errorln(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Warn(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Warnf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoinWithMessage(msg, args...))
}

func (l *TestLevelLogger) Warnln(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Info(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Infof(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoinWithMessage(msg, args...))
}

func (l *TestLevelLogger) Infoln(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Debug(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) Debugf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoinWithMessage(msg, args...))
}

func (l *TestLevelLogger) Debugln(args ...interface{}) {
	l.loggingData = append(l.loggingData, l.StringJoin(args...))
}

func (l *TestLevelLogger) StringJoinWithMessage(msg string, args ...interface{}) string {

	var message = fmt.Sprintf(msg, args...)
	return message
}

func (l *TestLevelLogger) StringJoin(args ...interface{}) string {

	var strs []string

	for _, value := range args {

		strs = append(strs, fmt.Sprintf("%v", value))
	}

	return strings.Join(strs, " ")
}
