package visitor_test

import (
	visitor2 "go.guide/patterns/visitor"
	"testing"
)

func TestVisitor(t *testing.T) {
	visitor := &visitor2.Visitor{}

	person := &visitor2.Person{Name: "aaa"}
	person.Accept(visitor)

	animal := &visitor2.Animal{Name: "bbb"}
	animal.Accept(visitor)
}
