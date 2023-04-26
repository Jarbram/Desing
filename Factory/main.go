package main

import (
	"factory/factory"
	"fmt"
	"os"
)

func main() {
	// Solicita al usuario el tipo de conexión a establecer
	var t int
	fmt.Print("write the connection that you want to do: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error in the connection: %v", err)
		os.Exit(1)
	}

	// Obtiene una instancia de la conexión utilizando una fábrica
	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("invalid connection")
		os.Exit(1)
	}

	// Se conecta a la base de datos
	err = conn.Connect()
	if err != nil {
		fmt.Printf("can't to do the connection: %v", err)
		os.Exit(1)
	}

	// Obtiene la hora actual del servidor y la imprime
	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("error in the connection: %v", err)
		os.Exit(1)
	}
	fmt.Println(now.Format("2006-01-02"))

	// Cierra la conexión
	err = conn.Close()
	if err != nil {
		fmt.Printf("can't close the connection: %v", err)
		os.Exit(1)
	}
}
