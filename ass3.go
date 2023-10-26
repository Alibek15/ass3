//                 .
//                .;;:,.
//                 ;iiii;:,.                                   .,:;.
//                 :i;iiiiii:,                            .,:;;iiii.
//                  ;iiiiiiiii;:.                    .,:;;iiiiii;i:
//                   :iiiiiiiiiii:......,,,,,.....,:;iiiiiiiiiiii;
//                    ,iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii:
//                     .:iii;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;,
//                       .:;;iiiiiiiiiiiiiiiiiiiiiiiiiii;;ii;,
//                        :iiii;;iiiiiiiiiiiiiii;;iiiiiii;:.
//                       ,iiii;1f:;iiiiiiiiiiii;if;:iiiiiii.
//                      .iiiii:iL..iiiiiiiiiiii;:f: iiiiiiii.
//                      ;iiiiii:.,;iiii;iiiiiiii:..:iiiiiiii:
//                     .i;;;iiiiiiiiii;,,;iiiiiiiiiiii;;iiiii.
//                     ::,,,,:iiiiiiiiiiiiiiiiiiiiii:,,,,:;ii:
//                     ;,,,,,:iiiiiiii;;;;;;;iiiiii;,,,,,,;iii.
//                     ;i;;;;iiiiiiii;:;;;;;:iiiiiii;::::;iiii:
//                     ,iiiiiiiiiiiiii;;;;;;:iiiiiiiiiiiiiiiiii.
//                      .iiiiiiiiiiiiii;;;;;iiiiiiiiiiiiiiiiiii:
//                       .;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;
//                        ;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii.
//                       .;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,

package main

import (
	"fmt"
	"sync"
)


type RemoteControl struct {
	commands map[string]Command
	mu       sync.RWMutex
}

var remote *RemoteControl
var once sync.Once

func GetRemoteControl() *RemoteControl {
	once.Do(func() {
		remote = &RemoteControl{
			commands: make(map[string]Command),
		}
	})
	return remote
}

type Command interface {
	Execute()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light}
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light}
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

type Light struct {
	name string
}

func NewLight(name string) *Light {
	return &Light{name}
}

func (l *Light) TurnOn() {
	fmt.Printf("%s light is turned on.\n", l.name)
}

func (l *Light) TurnOff() {
	fmt.Printf("%s light is turned off.\n", l.name)
}

func main() {
	remote := GetRemoteControl()
	light := NewLight("Living Room")

	lightOnCmd := NewLightOnCommand(light)
	lightOffCmd := NewLightOffCommand(light)

	remote.RegisterCommand("lightOn", lightOnCmd)
	remote.RegisterCommand("lightOff", lightOffCmd)

	
	remote.ExecuteCommand("lightOn")
	remote.ExecuteCommand("lightOff")
}
