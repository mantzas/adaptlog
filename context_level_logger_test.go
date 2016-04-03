package adaptlog

import (
	"testing"
)

type contextLevelLogTestCase struct {
	logFunctionName string
	msg             string
	args            []string
}

var contextLogTests = []struct {
	in  contextLevelLogTestCase
	out string
}{
	{contextLevelLogTestCase{"Debug", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
	{contextLevelLogTestCase{"Debugf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "TestContext args=arg1,arg2,arg3"},
	{contextLevelLogTestCase{"Debugln", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},

	{contextLevelLogTestCase{"Info", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
	{contextLevelLogTestCase{"Infof", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "TestContext args=arg1,arg2,arg3"},
	{contextLevelLogTestCase{"Infoln", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},

	{contextLevelLogTestCase{"Warn", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
	{contextLevelLogTestCase{"Warnf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "TestContext args=arg1,arg2,arg3"},
	{contextLevelLogTestCase{"Warnln", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},

	{contextLevelLogTestCase{"Error", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
	{contextLevelLogTestCase{"Errorf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "TestContext args=arg1,arg2,arg3"},
	{contextLevelLogTestCase{"Errorln", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},

	{contextLevelLogTestCase{"Fatal", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
	{contextLevelLogTestCase{"Fatalf", "args=%s,%s,%s", []string{"arg1", "arg2", "arg3"}}, "TestContext args=arg1,arg2,arg3"},
	{contextLevelLogTestCase{"Fatalln", "", []string{"arg1", "arg2", "arg3"}}, "TestContext arg1 arg2 arg3"},
}

type logFunction func(string, ...string)

func TestContextLeveledLogger(t *testing.T) {

	var logger = new(TestLevelLogger)
	ConfigureLevelLogger(logger)
	clogger := NewContextLevelLogger("TestContext")

	logFunctions := make(map[string]logFunction, 0)
	logFunctions["Debug"] = func(msg string, args ...string) { clogger.Debug(convertToInterfaceSlice(args)...) }
	logFunctions["Debugf"] = func(msg string, args ...string) { clogger.Debugf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Debugln"] = func(msg string, args ...string) { clogger.Debugln(convertToInterfaceSlice(args)...) }

	logFunctions["Info"] = func(msg string, args ...string) { clogger.Info(convertToInterfaceSlice(args)...) }
	logFunctions["Infof"] = func(msg string, args ...string) { clogger.Infof(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Infoln"] = func(msg string, args ...string) { clogger.Infoln(convertToInterfaceSlice(args)...) }

	logFunctions["Warn"] = func(msg string, args ...string) { clogger.Warn(convertToInterfaceSlice(args)...) }
	logFunctions["Warnf"] = func(msg string, args ...string) { clogger.Warnf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Warnln"] = func(msg string, args ...string) { clogger.Warnln(convertToInterfaceSlice(args)...) }

	logFunctions["Error"] = func(msg string, args ...string) { clogger.Error(convertToInterfaceSlice(args)...) }
	logFunctions["Errorf"] = func(msg string, args ...string) { clogger.Errorf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Errorln"] = func(msg string, args ...string) { clogger.Errorln(convertToInterfaceSlice(args)...) }

	logFunctions["Fatal"] = func(msg string, args ...string) { clogger.Fatal(convertToInterfaceSlice(args)...) }
	logFunctions["Fatalf"] = func(msg string, args ...string) { clogger.Fatalf(msg, convertToInterfaceSlice(args)...) }
	logFunctions["Fatalln"] = func(msg string, args ...string) { clogger.Fatalln(convertToInterfaceSlice(args)...) }

	for _, testCase := range contextLogTests {
		logger.loggingData = nil
		logFunctions[testCase.in.logFunctionName](testCase.in.msg, testCase.in.args...)

		if logger.loggingData[0] != testCase.out {
			t.Fatalf("Expected [%s] Actual [%s]", testCase.out, logger.loggingData[0])
			t.FailNow()
		}
	}
}

func convertToInterfaceSlice(a []string) []interface{} {
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	return b
}
