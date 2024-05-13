package command

import "testing"

func TestForCommand(t *testing.T) {
	light := &Light{}
	switcher := &Switch{
		OnCommand:  &LightOnCommand{light: light},
		OffCommand: &LightOffCommand{light: light},
	}

	switcher.On()
	switcher.Off()
}
