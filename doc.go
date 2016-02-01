/*
Package adaptlog is a logging abstraction for go. The name of the package is a composition of adapt(adaptive) and log(logging).

The developer uses this abstraction in order to avoid depending on a specific logging implementation. The package provides two abstractions:

The standard logger based on the standard log package and the leveled logger based on many leveled logging packages out there.

The developer has only to setup either one by implementing a interface.

The simplest way to use adaptlog's standard logger is by simply implementing the StdLogger interface like below.

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

  func main() {

    // configure once
    adaptlog.ConfigStandardLogger(new(MyLogger))

    // get new standard logger
    logger, err := adaptlog.NewStandardLogger()

    if err != nil {
         panic(err.Error())
    }

    // use logger
    logger.Print("Hello World!")
  }


For a full guide visit https://github.com/mantzas/adaptlog
*/
package adaptlog
