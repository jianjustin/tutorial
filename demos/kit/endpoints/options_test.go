package endpoints

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

type Options struct {
	Timeout   time.Duration
	Retries   int
	Verbose   bool
	LogOutput io.Writer
}

type ConfigFunc func(*Options)

func WithTimeout(timeout time.Duration) ConfigFunc {
	return func(options *Options) {
		options.Timeout = timeout
	}
}

func WithRetries(retries int) ConfigFunc {
	return func(options *Options) {
		options.Retries = retries
	}
}

func WithVerbose(verbose bool) ConfigFunc {
	return func(options *Options) {
		options.Verbose = verbose
	}
}

func WithLogOutput(logOutput io.Writer) ConfigFunc {
	return func(options *Options) {
		options.LogOutput = logOutput
	}
}

func TestForFunction(t *testing.T) {
	options := &Options{
		Timeout:   time.Second * 5,
		Retries:   3,
		Verbose:   false,
		LogOutput: os.Stdout,
	}

	WithVerbose(true)(options)
	WithTimeout(time.Second * 10)(options)
	WithRetries(5)(options)
	WithLogOutput(os.Stderr)(options)

	// 使用配置好的选项执行具体的操作
	fmt.Println("Options:", options)
}
