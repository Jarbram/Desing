package singleton

import "sync"

type person struct {
	name string
	age  int
	mux  sync.Mutex // Mutex para garantizar exclusión mutua
}

var (
	p    *person   // Puntero a instancia única
	once sync.Once // Variable para garantizar que la instancia solo se cree una vez
)

// Función que devuelve la instancia única de "person"
func GetInstance() *person {
	once.Do(func() { // Solo se ejecuta una vez
		p = &person{} // Creación de la instancia
	})
	return p
}

// Función para establecer el nombre de la persona

func (p *person) SetName(n string) {
	p.mux.Lock()   // Bloqueo exclusivo en la estructura
	p.name = n     // Actualización del valor
	p.mux.Unlock() // Desbloqueo
}

// Función para obtener el nombre de la persona
func (p *person) GetName() string {
	p.mux.Lock()         // Bloqueo exclusivo en la estructura
	defer p.mux.Unlock() // Se asegura el desbloqueo al final de la función
	return p.name        // Devuelve el valor
}

// Función para incrementar en uno la edad de la persona
func (p *person) SetAge() {
	p.mux.Lock()   // Bloqueo exclusivo en la estructura
	p.age++        // Incremento del valor
	p.mux.Unlock() // Desbloqueo
}

// Función para obtener la edad de la persona
func (p *person) GetAge() int {
	p.mux.Lock()         // Bloqueo exclusivo en la estructura
	defer p.mux.Unlock() // Se asegura el desbloqueo al final de la función
	return p.age         // Devuelve el valor
}
