package main

import (
	"fmt"
	"strings"
)

func main() {
	country := "Japan"
	for i, c := range country {
		//d:pos c:car
		fmt.Printf("%d:%q", i, c)
	}
	fmt.Println()
	fmt.Println(country[1:3]) // retorna la variable country desde el indice 1 al 3-1 ( retorna ap)
	fmt.Println(country[:4])  // retorna la variable country desde el indice 0 al 4-1 ( retorna Japa)
	fmt.Println(country[2:])  // retorna la variable country desde el indice 2 al final ( retorna pan)

	fmt.Println(strings.Contains("Japan", "a")) // retorna si la variable country contiene el string a (retorna true)

	fmt.Println(strings.Count("Japan", "a")) // retorna cuántas veces tien la variable country el string a (retorna 2)

	fmt.Println("Japan" < "japan") //retorna si el string Japan es menor que japan ( retorna true)
}
