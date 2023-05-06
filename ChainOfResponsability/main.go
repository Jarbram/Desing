package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interfaz que define los métodos necesarios para la cadena de procesamiento
type PurchaseHandler interface {
	SetNext(handler PurchaseHandler)
	Process(purchase int) bool
}

// Objeto procesador base
type BasePurchaseHandler struct {
	nextHandler PurchaseHandler
}

func (b *BasePurchaseHandler) SetNext(handler PurchaseHandler) {
	b.nextHandler = handler
}

// Objeto procesador para compras pequeñas
type SmallPurchaseHandler struct {
	BasePurchaseHandler
}

func (s *SmallPurchaseHandler) Process(purchase int) bool {
	if purchase < 100 {
		fmt.Println("Compra pequeña aceptada")
		return true
	} else if s.nextHandler != nil {
		return s.nextHandler.Process(purchase)
	}
	return false
}

// Objeto procesador para compras medianas
type MediumPurchaseHandler struct {
	BasePurchaseHandler
}

func (m *MediumPurchaseHandler) Process(purchase int) bool {
	if purchase < 1000 {
		fmt.Println("Compra mediana aceptada")
		return true
	} else if m.nextHandler != nil {
		return m.nextHandler.Process(purchase)
	}
	return false
}

// Objeto procesador para compras grandes
type LargePurchaseHandler struct {
	BasePurchaseHandler
}

func (l *LargePurchaseHandler) Process(purchase int) bool {
	if purchase < 10000 {
		fmt.Println("Compra grande aceptada")
		return true
	} else if l.nextHandler != nil {
		return l.nextHandler.Process(purchase)
	}
	return false
}

func main() {
	// Crear objetos procesadores
	smallHandler := &SmallPurchaseHandler{}
	mediumHandler := &MediumPurchaseHandler{}
	largeHandler := &LargePurchaseHandler{}

	// Encadenar objetos procesadores
	smallHandler.SetNext(mediumHandler)
	mediumHandler.SetNext(largeHandler)

	// Ciclo de entrada de consola
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el monto de compra o 'q' para salir: ")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "q" {
			break
		}
		purchase, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		smallHandler.Process(purchase)
		fmt.Println("Ingrese el monto de compra o 'q' para salir: ")
	}
	fmt.Println("Saliendo...")
}
