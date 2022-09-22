package log

import (
	"github.com/rs/zerolog"
	"path"
	"runtime"
)

type callerHook struct {
}

func (h callerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if _, file, line, ok := runtime.Caller(5); ok {
		e.Str("file", path.Base(file)).Int("line", line)
	}
	if level == zerolog.ErrorLevel {
		// send notice
	}
}
