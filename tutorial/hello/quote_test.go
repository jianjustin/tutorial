//go:build ignore

package hello

import (
	"fmt"
	"rsc.io/quote"
	"testing"
)

func TestForQuote(t *testing.T) {
	fmt.Println(quote.Go())
}
