# adaptlog  [![alt text](https://godoc.org/github.com/mantzas/adaptlog?status.png)](https://godoc.org/github.com/mantzas/adaptlog)&nbsp;[![build status](https://img.shields.io/travis/mantzas/adaptlog.svg)](http://travis-ci.org/mantzas/adaptlog)&nbsp;[![Coverage Status](https://coveralls.io/repos/github/mantzas/adaptlog/badge.svg?branch=master)](https://coveralls.io/github/mantzas/adaptlog?branch=master)


Package adaptlog is a logging abstraction for go. The name of the package is a composition of adapt(adaptive) and log(logging).

The developer uses this abstraction in order to avoid depending on a specific logging implementation. 
The package provides a abstraction that covers the standard logging(like in the standard log package) and the leveled logging (like many of them out there).

The simplest way to use adaptlog's logger is by simply implementing the Logger interface like illustarted in the samples of the examples folder.

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
    ...

    func main() {

      // configure once
	  adaptlog.Configure(new(MyLogger), adaptlog.AnyLevel)

	  // use logger
	  adaptlog.Print("Hello World!")
    }
