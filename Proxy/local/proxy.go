package local

import "proxy/book"

// La interfaz Proxy define los métodos a ser implementados por un proxy.
type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
