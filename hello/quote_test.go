package main

import (
	"fmt"
	"testing"

	"rsc.io/quote"
)

func TestQuoteGo(t *testing.T) {
	fmt.Println(quote.Go())
}
