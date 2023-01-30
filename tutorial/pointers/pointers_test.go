package pointers

import (
	"testing"
)

func TestZeroval(t *testing.T) {
	i := 1

	zeroval(i)
	if i != 1 {
		t.Error("zeroval error")
	}
}

func TestZeroptr(t *testing.T) {
	i := 1

	zeroptr(&i)
	if i != 0 {
		t.Error("zeroptr error")
	}
}
