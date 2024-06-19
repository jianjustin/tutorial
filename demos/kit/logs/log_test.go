package logs_test

import (
	"bytes"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/log"
	stdlog "log"
	"os"
	"testing"
	"time"
)

func TestForDebugInfo(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)

	// make time predictable for this test
	baseTime := time.Date(2015, time.February, 3, 10, 0, 0, 0, time.UTC)
	mockTime := func() time.Time {
		baseTime = baseTime.Add(time.Second)
		return baseTime
	}

	logger = log.With(logger, "time", log.Timestamp(mockTime), "caller", log.DefaultCaller)

	logger.Log("call", "first")
	logger.Log("call", "second")
}

func TestForStructured(t *testing.T) {
	// Unstructured
	stdlog.Printf("HTTP server listening on %s", ":8080")

	// Structured
	logger := log.NewJSONLogger(os.Stdout)
	log.With(logger).Log("transport", "HTTP", "addr", ":8080", "msg", "listening")
}

func TestForContextual(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "instance_id", 123)

	logger.Log("msg", "starting")
	log.With(logger, "component", "worker").Log("status", "starting")
	log.With(logger, "component", "slacker").Log("status", "starting")
}

func TestStdlibWriterWithStd(t *testing.T) {
	//buf := &bytes.Buffer{}
	stdlog.SetOutput(os.Stderr)
	stdlog.SetFlags(stdlog.LstdFlags)
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	logger.Log("key", "val")
}

func TestStdlibWriterWithBuf(t *testing.T) {
	buf := &bytes.Buffer{}
	stdlog.SetOutput(buf)
	stdlog.SetFlags(stdlog.LstdFlags)
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	logger.Log("key", "val")
	fmt.Fprint(os.Stderr, buf.String())
}

func TestForStdlib(t *testing.T) {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger, kitlog.MessageKey("msg"), kitlog.TimestampKey("time")))
	stdlog.Print("I sure like pie")
}
