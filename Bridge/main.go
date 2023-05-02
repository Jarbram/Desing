package main

import "fmt"

// Abstracción
type Forma interface {
	Dibujar()
}

// Implementación
type PlataformaDeDibujo interface {
	DibujarRectangulo(x1, y1, x2, y2 int)
	DibujarCirculo(x, y, r int)
}

// Implementación concreta 1
type PlataformaDeDibujoCLI struct{}

func (p *PlataformaDeDibujoCLI) DibujarRectangulo(x1, y1, x2, y2 int) {
	fmt.Printf("Dibujando rectángulo en CLI: [%d,%d] to [%d,%d]\n", x1, y1, x2, y2)
}

func (p *PlataformaDeDibujoCLI) DibujarCirculo(x, y, r int) {
	fmt.Printf("Dibujando círculo en CLI: centro=[%d,%d], radio=%d\n", x, y, r)
}

// Implementación concreta 2
type PlataformaDeDibujoGUI struct{}

func (p *PlataformaDeDibujoGUI) DibujarRectangulo(x1, y1, x2, y2 int) {
	fmt.Printf("Dibujando rectángulo en GUI: [%d,%d] to [%d,%d]\n", x1, y1, x2, y2)
}

func (p *PlataformaDeDibujoGUI) DibujarCirculo(x, y, r int) {
	fmt.Printf("Dibujando círculo en GUI: centro=[%d,%d], radio=%d\n", x, y, r)
}

// Abstracción refinada
type Rectangulo struct {
	x1, y1, x2, y2 int
	plataforma     PlataformaDeDibujo
}

func (r *Rectangulo) Dibujar() {
	r.plataforma.DibujarRectangulo(r.x1, r.y1, r.x2, r.y2)
}

// Abstracción refinada
type Circulo struct {
	x, y, r    int
	plataforma PlataformaDeDibujo
}

func (c *Circulo) Dibujar() {
	c.plataforma.DibujarCirculo(c.x, c.y, c.r)
}

func main() {
	cli := &PlataformaDeDibujoCLI{}
	gui := &PlataformaDeDibujoGUI{}

	rectangulo := &Rectangulo{x1: 0, y1: 0, x2: 10, y2: 20, plataforma: cli}
	rectangulo.Dibujar() // Dibujando rectángulo en CLI: [0,0] to [10,20]

	circulo := &Circulo{x: 5, y: 5, r: 10, plataforma: gui}
	circulo.Dibujar() // Dibujando círculo en GUI: centro=[5,5], radio=10
}

//Claro, con gusto. En este ejemplo, se utiliza el patrón Bridge para separar la abstracción de la //implementación, permitiendo que ambas puedan variar independientemente.

//Primero, se define la interfaz Forma, que es la abstracción principal. Esta interfaz define el //método Dibujar(), que representa la acción de dibujar una forma genérica.

//A continuación, se define la interfaz PlataformaDeDibujo, que es la implementación principal. Esta interfaz define dos métodos, DibujarRectangulo y DibujarCirculo, que representan las acciones específicas de dibujar un rectángulo y un círculo.

//Luego, se definen dos implementaciones concretas de la interfaz PlataformaDeDibujo: PlataformaDeDibujoCLI y PlataformaDeDibujoGUI. Ambas implementaciones concretas definen los métodos DibujarRectangulo y DibujarCirculo de manera distinta, pero ambas cumplen con la interfaz PlataformaDeDibujo.

//A continuación, se definen dos clases de abstracción refinada: Rectangulo y Circulo. Estas clases implementan la interfaz Forma y tienen una referencia a una implementación concreta de PlataformaDeDibujo.

//Finalmente, en la función main, se crean una instancia de PlataformaDeDibujoCLI y otra instancia de PlataformaDeDibujoGUI. Luego, se crean una instancia de Rectangulo y otra de Circulo, pasando cada una la instancia de PlataformaDeDibujo correspondiente. Al llamar al método Dibujar de cada objeto, se puede ver cómo la implementación del método es diferente según la plataforma utilizada.

//En resumen, el patrón Bridge permite separar la abstracción de la implementación, permitiendo que ambas puedan variar independientemente y facilitando la creación de código modular y fácil de mantener.
