package endpoints

import (
	"fmt"
	"testing"
)

type Data struct {
	Value int
}

// 定义一个接口类型，表示Pipeline中的操作
type Operation interface {
	Process(data *Data)
}

// 实现一个操作的结构体类型
type OperationFunc func(data *Data)

func (f OperationFunc) Process(data *Data) {
	f(data)
}

func AddOne(data *Data) {
	data.Value += 1
}

func Double(data *Data) {
	data.Value *= 2
}

func Square(data *Data) {
	data.Value *= data.Value
}

type Pipeline struct {
	operations []Operation
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) AddOperation(operation Operation) {
	p.operations = append(p.operations, operation)
}

func (p *Pipeline) Run(data *Data) {
	for _, operation := range p.operations {
		operation.Process(data)
	}
}

func TestForPipeline(t *testing.T) {
	pipeline := NewPipeline()
	pipeline.AddOperation(OperationFunc(AddOne))
	pipeline.AddOperation(OperationFunc(Double))
	pipeline.AddOperation(OperationFunc(Square))

	data := &Data{Value: 2}
	pipeline.Run(data)

	fmt.Println(data.Value) // 输出：36
}
