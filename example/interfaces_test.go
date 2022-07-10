package example

import (
	"math"
	"testing"
)

func TestRect(t *testing.T) {
	r := rect{weight: 3, height: 4}

	if r.perim() != 14 {
		t.Error("rect error")
	}
}

func TestCircle(t *testing.T) {
	c := circle{radius: 5}

	if c.perim() != 2*math.Pi*c.radius {
		t.Error("circle error")
	}
}
