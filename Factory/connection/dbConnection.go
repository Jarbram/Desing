package connection

import "time"

// DBConnection es una interfaz que define los métodos que deben implementar las conexiones a la base de datos.
type DBConnection interface {
	// Connect establece una conexión con la base de datos. Devuelve un error si no se puede conectar.
	Connect() error
	// GetNow devuelve la hora actual de la base de datos y un error si no se puede obtener.
	GetNow() (*time.Time, error)
	// Close cierra la conexión a la base de datos. Devuelve un error si no se puede cerrar.
	Close() error
}
