package adaptlog

import (
	"strings"
)

// ContextLevelLogger is a level logger with a context
type ContextLevelLogger struct {
	logger  LevelLogger
	context string
}

// NewContextLevelLogger creates a new level logger with a specific context
func NewContextLevelLogger(context string) LevelLogger {
	return &ContextLevelLogger{Level, context}
}

// Fatal logging
func (c *ContextLevelLogger) Fatal(args ...interface{}) {
	c.logger.Fatal(c.prependContextToArgs(args...)...)
}

// Fatalf logging
func (c *ContextLevelLogger) Fatalf(msg string, args ...interface{}) {
	c.logger.Fatalf(strings.Join([]string{c.context, msg}, " "), args...)
}

// Fatalln logging
func (c *ContextLevelLogger) Fatalln(args ...interface{}) {
	c.logger.Fatalln(c.prependContextToArgs(args...)...)
}

// Error logging
func (c *ContextLevelLogger) Error(args ...interface{}) {
	c.logger.Error(c.prependContextToArgs(args...)...)
}

// Errorf logging
func (c *ContextLevelLogger) Errorf(msg string, args ...interface{}) {
	c.logger.Errorf(strings.Join([]string{c.context, msg}, " "), args...)
}

// Errorln logging
func (c *ContextLevelLogger) Errorln(args ...interface{}) {
	c.logger.Errorln(c.prependContextToArgs(args...)...)
}

// Warn logging
func (c *ContextLevelLogger) Warn(args ...interface{}) {
	c.logger.Warn(c.prependContextToArgs(args...)...)
}

// Warnf logging
func (c *ContextLevelLogger) Warnf(msg string, args ...interface{}) {
	c.logger.Warnf(strings.Join([]string{c.context, msg}, " "), args...)
}

// Warnln logging
func (c *ContextLevelLogger) Warnln(args ...interface{}) {
	c.logger.Warnln(c.prependContextToArgs(args...)...)
}

// Info logging
func (c *ContextLevelLogger) Info(args ...interface{}) {
	c.logger.Info(c.prependContextToArgs(args...)...)
}

// Infof logging
func (c *ContextLevelLogger) Infof(msg string, args ...interface{}) {
	c.logger.Infof(strings.Join([]string{c.context, msg}, " "), args...)
}

// Infoln logging
func (c *ContextLevelLogger) Infoln(args ...interface{}) {
	c.logger.Infoln(c.prependContextToArgs(args...)...)
}

// Debug logging
func (c *ContextLevelLogger) Debug(args ...interface{}) {
	c.logger.Debug(c.prependContextToArgs(args...)...)
}

// Debugf logging
func (c *ContextLevelLogger) Debugf(msg string, args ...interface{}) {
	c.logger.Infof(strings.Join([]string{c.context, msg}, " "), args...)
}

// Debugln logging
func (c *ContextLevelLogger) Debugln(args ...interface{}) {
	c.logger.Debugln(c.prependContextToArgs(args...)...)
}

func (c *ContextLevelLogger) prependContextToArgs(args ...interface{}) []interface{} {
	var prependedArgs []interface{}
	prependedArgs = append(prependedArgs, c.context)
	prependedArgs = append(prependedArgs, args...)
	return prependedArgs
}
