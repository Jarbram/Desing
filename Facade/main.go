package main

import (
	"fmt"
	"math"
)

type MathFacade interface {
	CalculateTriangleArea(base, height float64) float64
	CalculateCircleArea(radius float64) float64
	CalculateMean(numbers []float64) float64
}

type mathFacade struct{}

func NewMathFacade() MathFacade {
	return &mathFacade{}
}

func (m *mathFacade) CalculateTriangleArea(base, height float64) float64 {
	return 0.5 * base * height
}

func (m *mathFacade) CalculateCircleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func (m *mathFacade) CalculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	return sum / float64(len(numbers))
}

func main() {
	// Crear una nueva instancia de MathFacade
	math := NewMathFacade()

	// Calcular el área de un triángulo
	triangleArea := math.CalculateTriangleArea(5, 10)
	fmt.Println("El área del triángulo es:", triangleArea)

	// Calcular el área de un círculo
	circleArea := math.CalculateCircleArea(3)
	fmt.Println("El área del círculo es:", circleArea)

	// Calcular la media de un conjunto de números
	numbers := []float64{1, 2, 3, 4, 5}
	mean := math.CalculateMean(numbers)
	fmt.Println("La media de los números es:", mean)
}

// En este ejemplo, la interfaz MathFacade proporciona una interfaz simplificada para los clientes que necesitan realizar operaciones matemáticas complejas. La implementación de la interfaz MathFacade se encuentra en la estructura mathFacade, que encapsula los subsistemas complejos de trigonometría, geometría y estadística. La función NewMathFacade se encarga de crear una instancia de mathFacade.
