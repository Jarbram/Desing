package bike

import "fmt"

type Bike struct {
	Brand string
	Color string
}

func (b *Bike) Go() {
	fmt.Println("Go!!")
}
