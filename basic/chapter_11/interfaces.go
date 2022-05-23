package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sql := new(Square)
	sql.side = 5

	var areaIntf Shaper
	areaIntf = sql
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
