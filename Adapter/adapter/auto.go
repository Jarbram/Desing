package adapter

import "adapter/auto"

type AutoAdapter struct {
	Auto *auto.Auto
}

func (a *AutoAdapter) Move() {
	a.Auto.TurnOn()
	a.Auto.SpeedUp()
}
