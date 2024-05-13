package command

import "fmt"

// Command 是一个执行操作的接口
type Command interface {
	Execute()
}

// Light 是一个接收者，它知道如何执行操作
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	fmt.Println("Light is off")
}

// LightOnCommand 是一个命令，它会调用接收者的方法
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// LightOffCommand 是另一个命令，它也会调用接收者的方法
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Switch 是一个请求者，它会调用命令
type Switch struct {
	OnCommand  Command
	OffCommand Command
}

func (s *Switch) On() {
	s.OnCommand.Execute()
}

func (s *Switch) Off() {
	s.OffCommand.Execute()
}

func main() {
	light := &Light{}
	switcher := &Switch{
		OnCommand:  &LightOnCommand{light: light},
		OffCommand: &LightOffCommand{light: light},
	}

	switcher.On()
	switcher.Off()
}
