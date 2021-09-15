package model

type Person struct {
	Id   int
	Name string
}

func NewPerson(i int, n string) Person { //genera la instancia de la estructura Person
	return Person{i, n}
}
