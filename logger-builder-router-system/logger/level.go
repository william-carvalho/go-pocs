package logger

import "strings"

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l Level) IsValid() bool {
	return l >= DebugLevel && l <= ErrorLevel
}

func ParseLevel(value string) (Level, bool) {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "DEBUG":
		return DebugLevel, true
	case "INFO":
		return InfoLevel, true
	case "WARN":
		return WarnLevel, true
	case "ERROR":
		return ErrorLevel, true
	default:
		return 0, false
	}
}
