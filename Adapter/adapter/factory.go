package adapter

import (
	"adapter/auto"
	"adapter/bike"
)

func Factory(s string) Adapter {
	switch s {
	case "auto":
		d := auto.Auto{}
		return &AutoAdapter{&d}

	case "bike":
		d := bike.Bike{}
		return &BikeAdapter{&d}
	}
	return nil
}
