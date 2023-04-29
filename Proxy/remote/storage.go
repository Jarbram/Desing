package remote

import (
	"proxy/book"
	"time"
)

// Define la estructura Data que contiene los detalles del servidor y la lista de libros.
type Data struct {
	books    book.Books
	server   string
	port     uint16
	user     string
	password string
}

// Función para crear una nueva instancia de Data.
func New(server string, port uint16, user string, password string) *Data {
	d := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
	d.load()
	return d
}

// Busca un libro en la lista por su ID.
func (d *Data) ById(ID uint) book.Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}
	return book.Book{}
}

// Devuelve todos los libros de la lista.
func (d *Data) All() book.Books {
	time.Sleep(4 * time.Second)
	return d.books
}

// Carga algunos libros en la lista.
func (d *Data) load() {
	d.books = make(book.Books, 0, 5)
	d.books = append(d.books,
		book.Book{ID: 1, Title: "The Little Go Book", Name: "little-go-book", Author: "Karl Seguin"},
		book.Book{ID: 2, Title: "An Introduction to Programming in Go", Name: "introduction-to-programming-in-go", Author: "Caleb Doxsey"},
		book.Book{ID: 3, Title: "Go Bootcamp", Name: "go-bootcamp", Author: "Matt Aimonetti"},
		book.Book{ID: 4, Title: "Learning Go", Name: "learning-go", Author: "Miek Gieben"},
		book.Book{ID: 5, Title: "Network Programming with Go", Name: "network-programming-with-go", Author: "Jan Newmarch"},
	)
}
