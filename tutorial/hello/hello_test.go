package hello

import (
	"fmt"
	"testing"

	"go.guide/tutorial/tools"
)

func TestForHello(t *testing.T) {
	fmt.Println(tools.ToUpper("Hello"))
}
