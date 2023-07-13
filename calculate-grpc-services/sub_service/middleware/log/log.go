package log

import (
	"github.com/go-kit/log"
	"io"
	"time"
)

func InitLogger(writer io.Writer) log.Logger {
	logger := log.NewJSONLogger(writer)
	logger = log.With(logger, "time", log.TimestampFormat(func() time.Time {
		return time.Now()
	}, "2006-01-02 15:04:05"))
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}
