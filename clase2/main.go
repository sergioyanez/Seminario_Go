package main

import (
	"fmt"

	"tudai2021.com/model"
)

func changeName(p *model.Person, n string) { //acá apunta al mismo lugar de memoria
	p.Name = n
}

func cambioNombre(p *model.Person, n string) { //acá apunta al mismo lugar de memoria
	p.Name = n
}
func main() {
	p := model.NewPerson(46, "Elva")
	fmt.Println(p)
	changeName(&p, "Maria") //p es una referencia al lugar de memoria
	fmt.Println(p)
	cambioNombre(&p, "Sergio")
}
