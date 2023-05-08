package flags

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

func TestForCustomValue(t *testing.T) {
	var cv customValue
	flag.CommandLine.Var(&cv, "10", "custom Value implementation")
	flag.Parse()
	fmt.Println(cv)
}

func TestForSleep(t *testing.T) {
	period := flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("%v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

func TestForTempconv(t *testing.T) {
	f := celsiusFlag{Celsius: 20}
	flag.Var(&f, "temp", "the temperature")
	flag.Parse()
	fmt.Println(f.String())
}
