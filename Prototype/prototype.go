package main

import "fmt"

type prototype struct {
	name    string
	age     int
	friends []string
	color   *string
	phones  map[string]string
}

// Método String que muestra los valores actuales de los campos de la estructura
func (p *prototype) String() string {
	return fmt.Sprintf(
		"Name: %s, Age: %d, Friends: %v, Color: %s, Phones: %v",
		p.name,
		p.age,
		p.friends,
		*p.color,
		p.phones,
	)
}

// Método Clone que crea una nueva instancia que contiene los mismos valores.
func (p *prototype) Clone() prototype {
	// Se crea un nuevo slice de amigos (friends) y se copia el contenido del slice original a este nuevo slice
	friends := make([]string, len(p.friends))
	copy(friends, p.friends)

	// Se crea una nueva variable de color y se le asigna el valor del puntero original a esta variable
	color := *p.color

	// Se crea un nuevo mapa de teléfonos (phones) y se copia el contenido del mapa original a este nuevo mapa
	phones := make(map[string]string)

	for k, v := range p.phones {
		phones[k] = v
	}

	return prototype{
		name:    p.name,
		age:     p.age,
		friends: friends,
		color:   &color,
		phones:  phones,
	}
}
