package log

import (
	"github.com/go-kit/log"
	"io"
	"time"
)

func InitLogger(writer io.Writer) log.Logger {
	logger := log.NewJSONLogger(writer)
	logger = log.With(logger, "@timestamp", log.TimestampFormat(func() time.Time {
		return time.Now()
	}, time.RFC822))
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}
