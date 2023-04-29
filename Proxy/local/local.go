package local

import (
	"proxy/book"
	"proxy/remote"
)

// Local es una estructura que contiene los datos remotos y la caché local
type Local struct {
	Remote *remote.Data // datos remotos
	cache  book.Books   // caché local
}

// New crea una nueva instancia de Local
func New() *Local {
	return &Local{
		Remote: remote.New("https://anyplace.com", 8080, "user", "password"), // se crean los datos remotos
		cache:  make(book.Books, 0),                                          // se inicializa la caché local
	}
}

// GetByID obtiene un libro por su ID
func (l *Local) GetByID(ID uint) book.Book {
	// se busca en la caché local
	for _, v := range l.cache {
		if v.ID == ID {
			return v
		}
	}
	// si no se encuentra en la caché local, se obtiene de los datos remotos y se agrega a la caché local
	b := l.Remote.ById(ID)
	l.cache = append(l.cache, b)
	return b
}

// GetAll obtiene todos los libros
func (l *Local) GetAll() book.Books {
	// se busca en la caché local
	for _, v := range l.cache {
		if v.ID == 0 {
			return l.cache
		}
	}
	// si no se encuentra en la caché local, se obtiene de los datos remotos y se agrega a la caché local
	b := l.Remote.All()
	l.cache = append(l.cache, b...)
	return b
}
