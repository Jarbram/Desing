package adapter

import "adapter/bike"

type BikeAdapter struct {
	Bici *bike.Bike
}

func (b *BikeAdapter) Move() {
	b.Bici.Go()
}
