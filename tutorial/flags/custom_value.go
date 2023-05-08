package flags

import (
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
