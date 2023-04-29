package book

// definición del tipo de datos "Book"
type Book struct {
	ID     uint   // campo ID de tipo entero sin signo
	Title  string // campo Title de tipo string
	Name   string // campo Name de tipo string
	Author string // campo Author de tipo string
}

// definición del tipo de datos "Books", que es un slice de objetos de tipo "Book"

type Books []Book
