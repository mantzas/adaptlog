# adaptlog  [![alt text](https://godoc.org/github.com/mantzas/adaptlog?status.png)](https://godoc.org/github.com/mantzas/adaptlog)&nbsp;[![build status](https://img.shields.io/travis/mantzas/adaptlog.svg)](http://travis-ci.org/mantzas/adaptlog)&nbsp;[![Coverage Status](https://coveralls.io/repos/github/mantzas/adaptlog/badge.svg?branch=master)](https://coveralls.io/github/mantzas/adaptlog?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/mantzas/adaptlog)](https://goreportcard.com/report/github.com/mantzas/adaptlog)


Package adaptlog is a logging abstraction in go. The name of the package is a composition of adapt(adaptive) and log(logging).

The developer uses this abstraction to avoid depending on a specific logging implementation.
The package provides an abstraction that covers the standard logging(like in the standard log package) and the leveled logging (like many of them out there).

The simplest way to use adaptlog's logger is by simply implementing a Logger interface like illustrated in the samples of the examples folder.

    package main

    import (
        "log"

        "github.com/mantzas/adaptlog"
    )

    // MyLogger custom logger implementing the Simple Logger interface
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
        adaptlog.ConfigureSimpleLogger(new(MyLogger))

        // use logger
        adaptlog.Simple.Print("Hello World!")
    }
