package main

import (
	"fmt"
	"singleton/client_one"
	"singleton/client_two"
	"singleton/singleton"
	"sync"
)

func main() {
	// Se utiliza WaitGroup para permitir el control de concurrencia
	wg := sync.WaitGroup{}
	wg.Add(200)

	// Se crean goroutines para aumentar la concurrencia en la ejecución del codigo
	for i := 0; i < 100; i++ {
		go func() {
			// Ejecuta el método SetAge() del singleton a través del cliente client_one
			client_one.SetAge()
			// La goroutine ha terminado su tarea, se informa al WaitGroup
			wg.Done()
		}()
		go func() {
			// Ejecuta el método SetAge() del singleton a través del cliente client_two
			client_two.SetAge()
			// La goroutine ha terminado su tarea, se informa al WaitGroup
			wg.Done()
		}()
	}

	// Espera a que todas las goroutines concluyan su ejecución
	wg.Wait()

	// Obtengo la instancia única del singleton antes creado
	p := singleton.GetInstance()
	// Obtengo la edad establecida aleatoriamente mediante el método SetAge()
	age := p.GetAge()

	// Imprimo resultado por consola
	fmt.Println(age)
}
