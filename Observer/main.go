package main

import (
	"fmt"
)

// Observador
type User struct {
	name string
}

func (u *User) Update(carState string) {
	fmt.Printf("Hola %s, el estado del auto ha cambiado a %s\n", u.name, carState)
}

// Observado
type Car struct {
	state     string
	observers []Observer
}

func (c *Car) RegisterObserver(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *Car) SetState(state string) {
	c.state = state
	c.NotifyObservers()
}

func (c *Car) NotifyObservers() {
	for _, o := range c.observers {
		o.Update(c.state)
	}
}

// Interfaz que define los métodos necesarios para la clase observada y los observadores
type Observer interface {
	Update(carState string)
}

func main() {
	// Crear objeto observado
	car := &Car{}

	// Crear objetos observadores
	user1 := &User{name: "Juan"}
	user2 := &User{name: "Maria"}

	// Registrar observadores en objeto observado
	car.RegisterObserver(user1)
	car.RegisterObserver(user2)

	// Simular cambio de estado del auto
	car.SetState("En taller")
	car.SetState("Listo para recoger")
}
