package logging

import (
	"fmt"
	"log"
	"runtime"
)

type Logger interface {
	Debug(template string, args ...interface{})
	Info(template string, args ...interface{})
	Warning(template string, args ...interface{})
	Error(template string, args ...interface{})
}

var (
	nopLogger    = LogFunc(func(level LogLevel, template string, args ...interface{}) {})
	simpleLogger = LogFunc(func(level LogLevel, template string, args ...interface{}) {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = -1
		}
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
		log.Print(fmt.Sprintf("[%s:%d][%s] ", file, line, level.String()) + fmt.Sprintf(template, args...))
	})
)

// NewNopLogger creates a Logger that do nothing.
func NewNopLogger() Logger {
	return nopLogger
}

// NewSimpleLogger creates a Logger using the log package in stdlib.
func NewSimpleLogger() Logger {
	return simpleLogger
}
