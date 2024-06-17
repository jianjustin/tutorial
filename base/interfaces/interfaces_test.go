package interfaces

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

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	weight, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.weight * r.height
}

func (r rect) perim() float64 {
	return 2*r.weight + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
