package main

import (
	"fmt"
)

func main() {
	// Creación del prototipo original
	color := "red"
	phones := map[string]string{"home": "123-456-7890", "work": "123-456-7891"}
	p1 := prototype{"Abraham", 23, []string{"John", "George"}, &color, phones}
	// Creación de una copia del prototipo y modificación de algunos valores
	p2 := p1.Clone()
	p2.age = 40
	p2.name = "Lincoln"
	p2.friends[0] = "Mary"
	p2.friends[1] = "Todd"
	color = "blue"
	p2.phones["home"] = "123-456-7892"
	// Impresión de los prototipos original y modificado en formato String
	fmt.Println(p1.String())
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
