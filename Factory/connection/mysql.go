package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Bajo guión indica que no se usará directamente, pero que necesitamos dicho módulo para establecer conexión con MySQL
)

// Definimos el struct MySQL, que contendrá una conexión de base de datos.
type MySQL struct {
	db *sql.DB
}

// La función Connect establece la conexión con la base de datos MySQL.
func (m *MySQL) Connect() error {
	// El parámetro 'm' hace referencia a la instancia de MySQL que se ha creado

	dsn := "my-user:my-password@tcp(127.0.0.1:3306)/my-db"

	db, err := sql.Open("mysql", dsn) // Abrimos conexión con MySQL
	if err != nil {
		return err // Si ocurre un error, devuelve el mismo para su tratamiento
	}

	err = db.Ping() // Preguntamos a la base de datos si está disponible usando Ping()
	if err != nil { // Si ocurre un error en la conexión, éste será tratado como un error fatal.
		panic(err.Error())
	}

	m.db = db  // Establecemos la conexión de nuestro struct de MySQL con la base SQL
	return nil // Si todo sale bien, retorna nil para indicar que no ocurrió ningún error.
}

// GetNow es una función que retorna la hora actual del sistema.
func (m *MySQL) GetNow() (*time.Time, error) {
	var t time.Time
	var dateStr []uint8
	// Ejecutamos un comando SQL para obtener la fecha actual del sistema
	err := m.db.QueryRow("select CURRENT_DATE()").Scan(&dateStr)
	if err != nil {
		fmt.Printf("Error while getting current date: %s\n", err)
		return nil, err
	}

	dateString := string(dateStr)
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, dateString) // Convertimos la fecha en un objeto de tiempo tipo 'time.Time'
	if err != nil {
		fmt.Printf("Error while parsing date: %s\n", err)
		return nil, err
	}

	t = parsedTime // Guardamos la fecha en la variable t
	return &t, nil // Retornamos la dirección de memoria de la variable 't' y un puntero nulo para indicar que no hubo errores.
}

// Close se encarga de cerrar la conexión con la base de datos cuando deje de ser necesaria
func (m *MySQL) Close() error {
	return m.db.Close()
}
