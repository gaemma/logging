package logging

type LogLevel uint8

func (ll LogLevel) String() string {
	return levelMap[ll]
}

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

var levelMap = map[LogLevel]string{
	DebugLevel:   "DEBUG",
	InfoLevel:    "INFO",
	WarningLevel: "WARNING",
	ErrorLevel:   "ERROR",
}
