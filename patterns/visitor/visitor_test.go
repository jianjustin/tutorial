package visitor_test

import (
	"testing"

	visitor2 "go.guide/patterns/visitor"
)

func TestVisitor(t *testing.T) {
	var visitor visitor2.Visitor
	person := &visitor2.Person{Name: "aaa"}
	animal := &visitor2.Animal{Name: "bbb"}

	visitor = &visitor2.VisitorA{}
	person.Accept(visitor)
	animal.Accept(visitor)

	visitor = &visitor2.VisitorB{}
	person.Accept(visitor)
	animal.Accept(visitor)

	visitor = &visitor2.VisitorC{}
	person.Accept(visitor)
	animal.Accept(visitor)
}
