package logging

type LogFunc func(level LogLevel, template string, args ...interface{})

func (lf LogFunc) Debug(template string, args ...interface{}) {
	lf(DebugLevel, template, args...)
}

func (lf LogFunc) Info(template string, args ...interface{}) {
	lf(InfoLevel, template, args...)
}

func (lf LogFunc) Warning(template string, args ...interface{}) {
	lf(WarningLevel, template, args...)
}

func (lf LogFunc) Error(template string, args ...interface{}) {
	lf(ErrorLevel, template, args...)
}
