package main

import (
	"log"

	"github.com/mantzas/adaptlog"
)

// MyLogger custom logger implementing the StdLogger interface
type MyLogger struct {
}

// Print logging
func (l *MyLogger) Print(args ...interface{}) {
	log.Print(args...)
}

// Printf logging
func (l *MyLogger) Printf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Println logging
func (l *MyLogger) Println(args ...interface{}) {
	log.Println(args...)
}

// Panic logging
func (l *MyLogger) Panic(args ...interface{}) {
	log.Panic(args...)
}

// Panicf logging
func (l *MyLogger) Panicf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

// Panicln logging
func (l *MyLogger) Panicln(args ...interface{}) {
	log.Panicln(args...)
}

// Fatal logging
func (l *MyLogger) Fatal(args ...interface{}) {
	log.Panic(args...)
}

// Fatalf logging
func (l *MyLogger) Fatalf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

// Fatalln logging
func (l *MyLogger) Fatalln(args ...interface{}) {
	log.Panicln(args...)
}

// Error logging
func (l *MyLogger) Error(args ...interface{}) {
	log.Print(args...)
}

// Errorf logging
func (l *MyLogger) Errorf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Errorln logging
func (l *MyLogger) Errorln(args ...interface{}) {
	log.Println(args...)
}

// Warn logging
func (l *MyLogger) Warn(args ...interface{}) {
	log.Print(args...)
}

//Warnf logging
func (l *MyLogger) Warnf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Warnln logging
func (l *MyLogger) Warnln(args ...interface{}) {
	log.Println(args...)
}

// Info logging
func (l *MyLogger) Info(args ...interface{}) {
	log.Print(args...)
}

// Infof logging
func (l *MyLogger) Infof(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Infoln logging
func (l *MyLogger) Infoln(args ...interface{}) {
	log.Println(args...)
}

// Debug logging
func (l *MyLogger) Debug(args ...interface{}) {
	log.Print(args...)
}

// Debugf logging
func (l *MyLogger) Debugf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Debugln logging
func (l *MyLogger) Debugln(args ...interface{}) {
	log.Println(args...)
}

func main() {

	// configure once
	adaptlog.Configure(new(MyLogger), adaptlog.AnyLevel)

	// use logger
	adaptlog.Print("Hello World!")
}
