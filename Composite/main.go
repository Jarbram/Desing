package main

import "fmt"

// Definimos una interfaz común para los objetos individuales y los grupos de objetos
type Component interface {
	Execute() // Método común para ejecutar una operación en un objeto o grupo de objetos
}

// Definimos un objeto individual
type Leaf struct {
	name string // Atributo específico del objeto individual
}

// Implementamos el método Execute() para el objeto individual
func (l *Leaf) Execute() {
	fmt.Printf("Ejecutando operación en %s\n", l.name)
}

// Definimos un grupo de objetos
type Composite struct {
	components []Component // Slice de objetos individuales y grupos de objetos
}

// Método para añadir un objeto al grupo
func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

// Implementamos el método Execute() para el grupo
func (c *Composite) Execute() {
	fmt.Println("Ejecutando operación en grupo...")

	// Iteramos sobre el slice de objetos del grupo y ejecutamos el método Execute() en cada uno
	for _, component := range c.components {
		component.Execute()
	}
}

// Método para listar los objetos contenidos en el grupo
func (c *Composite) List() {
	fmt.Println("Lista de objetos en el grupo:")

	// Iteramos sobre el slice de objetos del grupo e imprimimos el nombre de cada uno
	for _, component := range c.components {
		fmt.Printf("- %s\n", component.(*Leaf).name)
	}
}

func main() {
	// Creamos tres objetos individuales
	leaf1 := &Leaf{name: "Hoja 1"}
	leaf2 := &Leaf{name: "Hoja 2"}
	leaf3 := &Leaf{name: "Hoja 3"}

	// Creamos un objeto Composite y añadimos las hojas al grupo
	composite := &Composite{}
	composite.Add(leaf1)
	composite.Add(leaf2)
	composite.Add(leaf3)

	// Ejecutamos la operación en cada hoja y en el grupo
	leaf1.Execute()
	leaf2.Execute()
	leaf3.Execute()
	composite.Execute()

	// Mostramos el contenido del grupo
	composite.List()

}
