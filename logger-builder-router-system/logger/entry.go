package logger

import "time"

type Fields map[string]any

func (f Fields) Clone() Fields {
	if len(f) == 0 {
		return nil
	}
	cloned := make(Fields, len(f))
	for key, value := range f {
		cloned[key] = value
	}
	return cloned
}

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     Level     `json:"level"`
	Message   string    `json:"message"`
	Fields    Fields    `json:"fields,omitempty"`
}
