package slices

import (
	"fmt"
	"testing"
)

func TestForMap(t *testing.T) {
	//创建字典
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("ken: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	//prs返回字典是否存在值标识
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

}
