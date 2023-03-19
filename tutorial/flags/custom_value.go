package main

import (
	"flag"
	"fmt"
	"strconv"
)

type customValue int

func (cv *customValue) String() string { return fmt.Sprintf("%v", *cv) }

func (cv *customValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*cv = customValue(v)
	return err
}

func main() {
	var cv customValue
	flag.CommandLine.Var(&cv, "10", "custom Value implementation")
	flag.Parse()
	fmt.Println(cv)
}
