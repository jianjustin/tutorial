package interfaces

import (
	"fmt"
	"testing"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func TestAnimals(t *testing.T) {
	dog := &Dog{}
	fmt.Println(dog.Speak())

	cat := &Cat{}
	fmt.Println(cat.Speak())
}
