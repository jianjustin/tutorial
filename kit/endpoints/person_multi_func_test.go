package endpoints

import (
	"fmt"
	"testing"
)

type Person struct {
	Name     string
	Age      int
	Language string
}

type PersonConfigFunc func(*Person)

func WithName(name string) PersonConfigFunc {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) PersonConfigFunc {
	return func(p *Person) {
		p.Age = age
	}
}

func WithLanguage(language string) PersonConfigFunc {
	return func(p *Person) {
		p.Language = language
	}
}

func NewPerson(configFuncs ...PersonConfigFunc) *Person {
	person := &Person{}

	for _, configFunc := range configFuncs {
		configFunc(person)
	}

	return person
}

func TestForPerson(t *testing.T) {
	person := NewPerson(
		WithName("John"),
		WithAge(30),
		WithLanguage("English"),
	)

	fmt.Println(person)
}
