package auto

import "fmt"

type Auto struct {
	Brand  string
	Model  string
	turnOn bool
}

func (a *Auto) TurnOn() {
	if a.turnOn {
		fmt.Println("Ya esta encendido")
	}
	fmt.Println("Encendiendo")
}

func (a *Auto) SpeedUp() {
	fmt.Println("Acelerando")
}
