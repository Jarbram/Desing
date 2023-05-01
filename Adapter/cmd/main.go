package main

import (
	"adapter/adapter"
	"fmt"
)

func main() {
	var t string
	fmt.Print("What type of vehicle? (auto/bike): ")
	fmt.Scanln(&t)

	transportAdapter := adapter.Factory(t)
	transportAdapter.Move()

}
