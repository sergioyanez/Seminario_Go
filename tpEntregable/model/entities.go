package model

import (
	"fmt"
	"unicode"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

type CadenaEntrante struct {
	Value string
}

func NewCadenaEntrante(value string) CadenaEntrante {
	return CadenaEntrante{value}
}

func GetTipo(cadena string) string {
	return cadena[:2]
}

func GetLength(cadena string) int {
	var i int
	if _, err := fmt.Sscanf(cadena[2:4], "%2d", &i); err == nil {

	}
	return i
}

func GetCantCaracteres(cadena string) int {
	return len(cadena[4:])
}

func GetValue(cadena string) string {
	return cadena[4:]
}

func NewResult(t string, l int, v string) Result {
	return Result{t, l, v}
}

func ChequeoDeCadena(tipo, valor string, largo int) bool {

	if tipo != "NN" && tipo != "TX" {
		return false
	}
	if tipo == "NN" && !esEntero(valor) {
		return false
	}
	if tipo == "TX" && !esLetra(valor) {
		return false
	}
	if largo != len(valor) {
		return false
	}
	return true

}

func esEntero(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func esLetra(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
