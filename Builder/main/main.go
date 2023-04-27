package main

import (
	builder "Builder"
	"fmt"
	"os"
)

func main() {
	var t int
	fmt.Print("Ingrese el tipo de mensaje que desea enviar: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer el tipo de mensaje : %v", err)
		os.Exit(1)
	}

	b := builder.Factory(t)
	if b == nil {
		fmt.Println("El tipo de mensaje no es válido")
		os.Exit(1)
	}
	sender := &builder.Sender{}
	sender.SetBuilder(b)
	msg, err := sender.BuildMessage("Gophers", "hello world")
	if err != nil {
		fmt.Printf("BuildMessage() have a error: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(msg.Body))
}

//importar el paquete builder
//capturar por consola que objeto desea si xml o json

//con el patron de diseno fabrica recibo el paquete que voy a construir y lo asigno al builder
