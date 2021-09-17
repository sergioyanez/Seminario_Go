package main

import (
	"fmt"
	"os"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

type Person struct { // este es un tipo Person que es una estructura
	Name string
	Age  int
}

type InterfacePerson interface {
	Describe(Logger) string
}

type Logger func(string) // este es un tipo que es una función

func (p *Person) Describe(l Logger) string {
	l("logger: " + p.Name)
	//	p.Name = "Martina Yañez"
	return p.Name
}

func main() {
	var p InterfacePerson
	p = &Person{"Cármen", 42}
	//	p2 := Person{Name: "Sergio", Age: 48}
	log := func(s string) {
		fmt.Println(s)
	}
	fmt.Println(p.Describe(log))
	//fmt.Println(p.Name)
	/*
		fmt.Println("Hola Mundo ")
		m := map[int]string{}
		m[0] = "Sergio"
		m[1] = " Yañez"
		for _, x := range m {
			fmt.Print(x)
		}
	*/
	/*
		for i := 0; i < 2; i++ {
			fmt.Print(m[i])
		}
	*/

	args := os.Args
	for _, x := range args {
		fmt.Println(x)
	}
	x := []int{}
	for i := 0; i < 5; i++ {
		x = append(x, i)
	}
	fmt.Println(x)

	/**	r, err := suma(5, 2)
	  	if err != nil {
	  		fmt.Println(err)
	  	} else {
	  		fmt.Println(r)
	  	}
	  	var arr [2]int
	  	arr[0] = 5
	  	arr[1] = 9
	  	for i, v := range arr {
	  		fmt.Println(i, v)
	  	}
	  	fmt.Println("------------------")
	  	var l []int
	  	l = append(l, 10)
	  	fmt.Printf("%p\n", l)
	  	l = append(l, 100)
	  	l = append(l, 1000)
	  	for i, v := range l {
	  		fmt.Println(i, v)
	  	}

	  }

	  func suma(a, b int) (int, error) {
	  	if a < b {
	  		return 0, errors.New("el primer valor a es < que el segundo valor  b")
	  	}
	  	return a + b, nil
	*/
}
