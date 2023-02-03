package strings

import (
	"fmt"
	"testing"
)

func TestForSPrintf(t *testing.T) {
	a := fmt.Sprintf("%d*%d=%d", 1, 2, 2)
	fmt.Println(a)
}
