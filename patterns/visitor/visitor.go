package visitor

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type Person struct {
	Name string
}

func (p *Person) Accept(v Visitor) {
	v.VisitPerson(p)
}

type Animal struct {
	Name string
}

func (a *Animal) Accept(v Visitor) {
	v.VisitAnimal(a)
}

type Visitor interface {
	VisitPerson(p *Person)
	VisitAnimal(a *Animal)
}

type VisitorA struct {
	Visitor
}

func (v *VisitorA) VisitPerson(p *Person) {
	fmt.Printf("Person %s A visited\n", p.Name)
}

func (v *VisitorA) VisitAnimal(a *Animal) {
	fmt.Printf("Animal %s A visited\n", a.Name)
}

type VisitorB struct {
	Visitor
}

func (v *VisitorB) VisitPerson(p *Person) {
	fmt.Printf("Person %s B visited\n", p.Name)
}

func (v *VisitorB) VisitAnimal(a *Animal) {
	fmt.Printf("Animal %s B visited\n", a.Name)
}

type VisitorC struct {
	VisitorB
}

func (v *VisitorC) VisitAnimal(a *Animal) {
	fmt.Printf("Animal %s C visited\n", a.Name)
}
