package structs_test

import (
	"fmt"
	"testing"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func TestForStructsFields(t *testing.T) {
	ms := struct1{}
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Println(ms)
}
