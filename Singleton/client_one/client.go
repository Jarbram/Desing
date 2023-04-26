// Paquete client_one contiene la función SetAge
package client_one

// Se importa el paquete singleton
import "singleton/singleton"

// Función SetAge llama el método SetAge de la instancia del singleton
func SetAge() {
	// Obtiene una instancia del objeto singleton
	p := singleton.GetInstance()
	// Llama al método SetAge en la instancia
	p.SetAge()
}
