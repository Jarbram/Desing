package main

import (
	"fmt"
	"proxy/local" // Importamos el paquete local que contiene la implementación de nuestro proxy.
	"time"
)

var loc local.Proxy // Declaramos una variable global para el objeto Proxy.

func main() {
	loc = local.New() // Inicializamos nuestro objeto Proxy.

	GetByID(2) // Llamamos a la función GetByID pasándole un ID como parámetro.
	GetByID(2)
	GetByID(2)
	GetByID(3)
	GetByID(4)
	GetByID(2)
	GetByID(5)
	GetAll() // Llamamos a la función GetAll.
	GetAll()
}

// GetByID busca en caché el objeto con el ID proporcionado.
func GetByID(ID uint) {
	start := time.Now()
	// Iniciamos un temporizador para medir el tiempo de ejecución.
	fmt.Printf("%+v", loc.GetByID(ID))
	// Imprimimos el resultado de la búsqueda.
	elasped := time.Since(start)
	// Detenemos el temporizador y calculamos la duración de la operación.
	fmt.Printf("Elapsed time: %s", elasped)
	// Imprimimos el tiempo de ejecución en segundos.
}

// GetAll obtiene todos los objetos almacenados en caché.
func GetAll() {
	start := time.Now()
	// Iniciamos un temporizador para medir el tiempo de ejecución.
	fmt.Printf("%+v", loc.GetAll())
	// Imprimimos el resultado de la búsqueda.
	elasped := time.Since(start)
	// Detenemos el temporizador y calculamos la duración de la operación.
	fmt.Printf("Elapsed time: %v", elasped)
	// Imprimimos el tiempo de ejecución en segundos.
}
