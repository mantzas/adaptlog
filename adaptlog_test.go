package adaptlog

import "testing"

const (
	Print   = "Print"
	Printf  = "Printf"
	Println = "Println"
)

type TestMinimumLogger struct {
	loggingData []string
}

func (l *TestMinimumLogger) Print(args ...interface{}) {
	l.loggingData = append(l.loggingData, Print)
}

func (l *TestMinimumLogger) Printf(msg string, args ...interface{}) {
	l.loggingData = append(l.loggingData, Printf)
}

func (l *TestMinimumLogger) Println(args ...interface{}) {
	l.loggingData = append(l.loggingData, Println)
}

func TestNewMinimumLoggerWithoutConfigReturnsError(t *testing.T) {

	logger, err := NewMinimumLogger()

	if logger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestNewMinimumLoggerSucceeds(t *testing.T) {

	var logger = new(TestMinimumLogger)

	ConfigMinimalLogger(logger)

	minLogger, err := NewMinimumLogger()

	if minLogger != nil {
		t.Fatal("Logger should have been nil!")
	}

	if err == nil {
		t.Fatal("Should have returned a error")
	}
}
