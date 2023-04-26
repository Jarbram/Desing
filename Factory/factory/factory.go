package factory

import "factory/connection"

// Factory crea una conexión a la base de datos según el tipo especificado
func Factory(t int) connection.DBConnection {
	switch t {
	// Si el tipo es 1, se retorna una conexión de tipo Postgres
	case 1:
		return &connection.Postgres{}
		// Si el tipo es 2, se retorna una conexión de tipo MySQL
	case 2:
		return &connection.MySQL{}
		// Si el tipo no corresponde a Postgres ni a MySQL, se retorna nil
	default:
		return nil
	}
}
