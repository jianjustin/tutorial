package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length float32
	width  float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println("looping through shapes for area ...")
	for _, v := range shapes {
		fmt.Println("Shape details: ", v)
		fmt.Println("Area of this shape is: ", v.Area())
	}
}
