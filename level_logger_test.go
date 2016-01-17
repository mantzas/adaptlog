package adaptlog

import "testing"

func TestNewLeveledLoggerWithoutConfigReturnsError(t *testing.T) {

	logger, err := NewLevelLogger()

	if logger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewLeveledLoggerSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigLevelLogger(logger)

	stdLogger, err := NewLevelLogger()

	if stdLogger == nil {
		t.Fatal("Logger should have been not nil!")
	}

	if err != nil {
		t.Fatal("Should not have returned a error")
	}
}

func TestNewLevelLoggerLoggingSucceeds(t *testing.T) {

	var logger = new(TestLevelLogger)

	ConfigLevelLogger(logger)

	stdLogger, err := NewLevelLogger()

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

	if logger.loggingData[0] != Fatal || logger.loggingData[1] != Fatalf || logger.loggingData[2] != Fatalln ||
		logger.loggingData[3] != Panic || logger.loggingData[4] != Panicf || logger.loggingData[5] != Panicln ||
		logger.loggingData[6] != Error || logger.loggingData[7] != Errorf || logger.loggingData[8] != Errorln ||
		logger.loggingData[9] != Warn || logger.loggingData[10] != Warnf || logger.loggingData[11] != Warnln ||
		logger.loggingData[12] != Info || logger.loggingData[13] != Infof || logger.loggingData[14] != Infoln ||
		logger.loggingData[15] != Debug || logger.loggingData[16] != Debugf || logger.loggingData[17] != Debugln {
		t.Fatal("Logged items do not match!")
	}
}

const (
	Error   = "Error"
	Errorf  = "Errorf"
	Errorln = "Errorln"

	Warn   = "Warn"
	Warnf  = "Warnf"
	Warnln = "Warnln"

	Info   = "Info"
	Infof  = "Infof"
	Infoln = "Infoln"

	Debug   = "Debug"
	Debugf  = "Debugf"
	Debugln = "Debugln"
)

type TestLevelLogger struct {
	loggingData []string
}

func (l *TestLevelLogger) Panic(args ...interface{}) {
	l.loggingData = append(l.loggingData, Panic)
}

func (l *TestLevelLogger) Panicf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Panicf)
}

func (l *TestLevelLogger) Panicln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Panicln)
}

func (l *TestLevelLogger) Fatal(args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatal)
}

func (l *TestLevelLogger) Fatalf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatalf)
}

func (l *TestLevelLogger) Fatalln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Fatalln)
}

func (l *TestLevelLogger) Error(args ...interface{}) {
	l.loggingData = append(l.loggingData, Error)
}

func (l *TestLevelLogger) Errorf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Errorf)
}

func (l *TestLevelLogger) Errorln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Errorln)
}

func (l *TestLevelLogger) Warn(args ...interface{}) {
	l.loggingData = append(l.loggingData, Warn)
}

func (l *TestLevelLogger) Warnf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Warnf)
}

func (l *TestLevelLogger) Warnln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Warnln)
}

func (l *TestLevelLogger) Info(args ...interface{}) {
	l.loggingData = append(l.loggingData, Info)
}

func (l *TestLevelLogger) Infof(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Infof)
}

func (l *TestLevelLogger) Infoln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Infoln)
}

func (l *TestLevelLogger) Debug(args ...interface{}) {
	l.loggingData = append(l.loggingData, Debug)
}

func (l *TestLevelLogger) Debugf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Debugf)
}

func (l *TestLevelLogger) Debugln(args ...interface{}) {
	l.loggingData = append(l.loggingData, Debugln)
}
