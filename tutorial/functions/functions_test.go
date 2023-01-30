package functions

import (
	"testing"
)

func TestPlusFunction(t *testing.T) {
	res := plus(1, 2)
	if res != 3 {
		t.Error("plus function error")
	}
}

func TestValsFunction(t *testing.T) {
	a, b := vals()
	if a != 3 || b != 7 {
		t.Error("vals function error")
	}
}

func TestSumFunction(t *testing.T) {
	total := sum(1, 2, 3, 4, 5)
	if total != 15 {
		t.Error("sum function error")
	}
}

func TestIntSeqFunction(t *testing.T) {
	nextInt := intSeq()
	if nextInt() != 1 {
		t.Error("intseq function error")
	}
	if nextInt() != 2 {
		t.Error("intseq function error")
	}
}
