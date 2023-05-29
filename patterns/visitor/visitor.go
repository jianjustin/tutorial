package visitor

import "fmt"

type Element interface {
	Accept(v *Visitor)
}

type Person struct {
	Name string
}

func (p *Person) Accept(v *Visitor) {
	v.VisitPerson(p)
}

type Animal struct {
	Name string
}

func (a *Animal) Accept(v *Visitor) {
	v.VisitAnimal(a)
}

type Visitor struct{}

func (v *Visitor) VisitPerson(p *Person) {
	fmt.Printf("Person %s visited\n", p.Name)
}

func (v *Visitor) VisitAnimal(a *Animal) {
	fmt.Printf("Animal %s visited\n", a.Name)
}
