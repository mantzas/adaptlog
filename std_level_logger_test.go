package adaptlog

import (
	"strings"
	"testing"
)

func TestLogLevelToString(t *testing.T) {
	if DebugLevel.String() != "Debug" {
		t.Fatal("Should have been Debug")
	}
	if InfoLevel.String() != "Info" {
		t.Fatal("Should have been Info")
	}
	if WarnLevel.String() != "Warn" {
		t.Fatal("Should have been Warn")
	}
	if ErrorLevel.String() != "Error" {
		t.Fatal("Should have been Error")
	}
	if FatalLevel.String() != "Fatal" {
		t.Fatal("Should have been Fatal")
	}

	if LogLevel(9).String() != "Not mapped value 9" {
		t.Fatal("Should have been 'Not mapped value 9'")
	}
}

func TestStdLevelLoggerNew(t *testing.T) {

	ConfigureStdLevelLogger(DebugLevel, nil)
	l := NewStdLevelLogger()

	if l == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestStdLevelLoggerNewWithWriter(t *testing.T) {

	ConfigureStdLevelLogger(DebugLevel, new(TestWriter))
	l := NewStdLevelLogger()

	if l == nil {
		t.Fatal("Should have returned a error")
	}
}

func TestStdLevelLoggerWithHigherDefaultLevel(t *testing.T) {

	w := new(TestWriter)

	ConfigureStdLevelLogger(ErrorLevel, w)
	l := NewStdLevelLogger()

	l.Warn("Test")
	l.Error("Test")
	l.Fatal("Test")

	if len(w.data) != 2 {
		t.Fatalf("Data should have been 2 %s", w.data)
	}
	if strings.HasSuffix(w.data[0], "Error Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}

	if strings.HasSuffix(w.data[1], "Fatal Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}
	w.data = nil

	l.Warnf("Test")
	l.Errorf("Test")
	l.Fatalf("Test")

	if len(w.data) != 2 {
		t.Fatalf("Data should have been 2 %s", w.data)
	}
	if strings.HasSuffix(w.data[0], "Error Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}

	if strings.HasSuffix(w.data[1], "Fatal Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}
	w.data = nil

	l.Warnln("Test")
	l.Errorln("Test")
	l.Fatalln("Test")

	if len(w.data) != 2 {
		t.Fatalf("Data should have been 2 %s", w.data)
	}
	if strings.HasSuffix(w.data[0], "Error Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}

	if strings.HasSuffix(w.data[1], "Fatal Test") {
		t.Fatalf("Expected %s actual %s", "Error Test", w.data[0])
	}
}

var stdLevelLogTests = []struct {
	in  contextLevelLogTestCase
	out string
}{
	{contextLevelLogTestCase{"Debug", "", []string{"arg1", "arg2", "arg3"}}, "Debugarg1arg2arg3\n"},
	{contextLevelLogTestCase{"Debugf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "Debug args=arg1,arg2,arg3\n"},
	{contextLevelLogTestCase{"Debugln", "", []string{"arg1", "arg2", "arg3"}}, "Debug arg1 arg2 arg3\n"},

	{contextLevelLogTestCase{"Info", "", []string{"arg1", "arg2", "arg3"}}, "Infoarg1arg2arg3\n"},
	{contextLevelLogTestCase{"Infof", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "Info args=arg1,arg2,arg3\n"},
	{contextLevelLogTestCase{"Infoln", "", []string{"arg1", "arg2", "arg3"}}, "Info arg1 arg2 arg3\n"},

	{contextLevelLogTestCase{"Warn", "", []string{"arg1", "arg2", "arg3"}}, "Warnarg1arg2arg3\n"},
	{contextLevelLogTestCase{"Warnf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "Warn args=arg1,arg2,arg3\n"},
	{contextLevelLogTestCase{"Warnln", "", []string{"arg1", "arg2", "arg3"}}, "Warn arg1 arg2 arg3\n"},

	{contextLevelLogTestCase{"Error", "", []string{"arg1", "arg2", "arg3"}}, "Errorarg1arg2arg3\n"},
	{contextLevelLogTestCase{"Errorf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "Error args=arg1,arg2,arg3\n"},
	{contextLevelLogTestCase{"Errorln", "", []string{"arg1", "arg2", "arg3"}}, "Error arg1 arg2 arg3\n"},

	{contextLevelLogTestCase{"Fatal", "", []string{"arg1", "arg2", "arg3"}}, "Fatalarg1arg2arg3\n"},
	{contextLevelLogTestCase{"Fatalf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "Fatal args=arg1,arg2,arg3\n"},
	{contextLevelLogTestCase{"Fatalln", "", []string{"arg1", "arg2", "arg3"}}, "Fatal arg1 arg2 arg3\n"},
}

func TestStdLeveledLogger(t *testing.T) {

	w := new(TestWriter)
	ConfigureStdLevelLogger(DebugLevel, w)
	l := NewStdLevelLogger()

	logFunctions := make(map[string]logFunction, 0)
	logFunctions["Debug"] = func(msg string, args ...string) { l.Debug(convertToInterfaceSlice(args)...) }
	logFunctions["Debugf"] = func(msg string, args ...string) { l.Debugf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Debugln"] = func(msg string, args ...string) { l.Debugln(convertToInterfaceSlice(args)...) }

	logFunctions["Info"] = func(msg string, args ...string) { l.Info(convertToInterfaceSlice(args)...) }
	logFunctions["Infof"] = func(msg string, args ...string) { l.Infof(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Infoln"] = func(msg string, args ...string) { l.Infoln(convertToInterfaceSlice(args)...) }

	logFunctions["Warn"] = func(msg string, args ...string) { l.Warn(convertToInterfaceSlice(args)...) }
	logFunctions["Warnf"] = func(msg string, args ...string) { l.Warnf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Warnln"] = func(msg string, args ...string) { l.Warnln(convertToInterfaceSlice(args)...) }

	logFunctions["Error"] = func(msg string, args ...string) { l.Error(convertToInterfaceSlice(args)...) }
	logFunctions["Errorf"] = func(msg string, args ...string) { l.Errorf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Errorln"] = func(msg string, args ...string) { l.Errorln(convertToInterfaceSlice(args)...) }

	logFunctions["Fatal"] = func(msg string, args ...string) { l.Fatal(convertToInterfaceSlice(args)...) }
	logFunctions["Fatalf"] = func(msg string, args ...string) { l.Fatalf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Fatalln"] = func(msg string, args ...string) { l.Fatalln(convertToInterfaceSlice(args)...) }

	for _, testCase := range stdLevelLogTests {
		w.data = nil
		logFunctions[testCase.in.logFunctionName](testCase.in.msg, testCase.in.args...)

		if !strings.HasSuffix(w.data[0], testCase.out) {
			t.Fatalf("Case %s: Expected [%s] Actual [%s]", testCase.in.logFunctionName, testCase.out, w.data[0])
		}
	}
}

type TestWriter struct {
	data []string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	var text = string(p)
	t.data = append(t.data, text)
	return len(text), nil
}
