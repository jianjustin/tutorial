package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{name: "Alice", age: 30})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
}
