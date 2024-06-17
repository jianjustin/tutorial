package slices

import (
	"fmt"
	"testing"
)

func TestForMapsForrange(t *testing.T) {
	map1 := make(map[int]float32)

	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is %d - value is：%f\n", key, value)
	}
}
