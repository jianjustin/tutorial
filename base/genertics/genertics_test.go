package genertics

import (
	"fmt"
	"testing"
)

func Add[T int | float64](a, b T) T {
	return a + b
}

func TestForAdd(t *testing.T) {
	fmt.Println(Add(1, 2))     // 输出: 3
	fmt.Println(Add(1.5, 2.5)) // 输出: 4.0
}

// 泛型类型
type Container[T int | string] struct {
	value T
}

func (c *Container[T]) Set(value T) {
	c.value = value
}

func (c *Container[T]) Get() T {
	return c.value
}

func TestForContainer(t *testing.T) {
	intContainer := Container[int]{}
	intContainer.Set(42)
	fmt.Println(intContainer.Get()) // 输出: 42

	stringContainer := Container[string]{}
	stringContainer.Set("Hello")
	fmt.Println(stringContainer.Get()) // 输出: Hello
}

type Number interface {
	int | float64
}

func Multiply[T Number](a, b T) T {
	return a * b
}

func TestForMultiply(t *testing.T) {
	fmt.Println(Multiply(2, 3))     // 输出: 6
	fmt.Println(Multiply(2.5, 3.5)) // 输出: 8.75
}

type Pair[T any, U any] struct {
	first  T
	second U
}

func (p *Pair[T, U]) SetFirst(value T) {
	p.first = value
}

func (p *Pair[T, U]) SetSecond(value U) {
	p.second = value
}

func (p *Pair[T, U]) GetFirst() T {
	return p.first
}

func (p *Pair[T, U]) GetSecond() U {
	return p.second
}

func TestForPair(t *testing.T) {
	pair := Pair[int, string]{}
	pair.SetFirst(1)
	pair.SetSecond("one")
	fmt.Println(pair.GetFirst())  // 输出: 1
	fmt.Println(pair.GetSecond()) // 输出: one
}

type Adder[T any] interface {
	Add(a, b T) T
}

type IntAdder struct{}

func (IntAdder) Add(a, b int) int {
	return a + b
}

type FloatAdder struct{}

func (FloatAdder) Add(a, b float64) float64 {
	return a + b
}

func TestForAdder(t *testing.T) {
	var intAdder Adder[int] = IntAdder{}
	fmt.Println(intAdder.Add(1, 2)) // 输出: 3

	var floatAdder Adder[float64] = FloatAdder{}
	fmt.Println(floatAdder.Add(1.5, 2.5)) // 输出: 4.0
}
