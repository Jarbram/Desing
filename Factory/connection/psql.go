package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct { // Define una estructura llamada Postgres
	db *sql.DB // db es un puntero a sql.DB
}

func (p *Postgres) Connect() error { // Define un método Connect para la estructura Postgres
	dsn := "postgres://my-user:my-password@localhost:5432/my-db?sslmode=disable"

	db, err := sql.Open("postgres", dsn) // Crea la conexión usando el driver Postgres y la variable DSN
	if err != nil {
		return err // Si hay un error, lo retorna
	}
	p.db = db  // Asigna la conexión a la variable db de la estructura Postgres
	return nil // Si todo sale bien, retorna nil (sin errores)
}

func (p *Postgres) GetNow() (*time.Time, error) { // Define un método GetNow para la estructura Postgres
	t := &time.Time{}                                   // Crea un nuevo objeto Time
	err := p.db.QueryRow("select CURRENT_DATE").Scan(t) // Ejecuta la consulta y guarda el resultado en la variable t
	if err != nil {
		fmt.Printf("Error while getting current date: %s\n", err) // Si hay un error, lo imprime en consola
		return nil, err                                           // Retorna el error
	}
	return t, nil // Si todo sale bien, retorna el objeto Time y nil (sin errores)
}

func (p *Postgres) Close() error { // Define un método Close para la estructura Postgres
	return p.db.Close() // Cierra la conexión a la base de datos y retorna nil (sin errores)
}
