package main

import (
	"entregableGoTudai2021/model"
	"errors"
	"fmt"
)

func generarResultado(cad string) (model.Result, error) {
	if len(cad) > 4 {
		cadena := model.NewCadenaEntrante(cad)
		tipo := model.GetTipo(cadena.Value)
		valor := model.GetValue(cadena.Value)
		largo := model.GetLength(cadena.Value)
		if model.ChequeoDeCadena(tipo, valor, largo) {
			return model.NewResult(tipo, largo, valor), nil

		}
	}
	return model.Result{}, errors.New("cadena Invalida")

}

func main() {
	nuevaCadena := "NN03897"

	result, err := generarResultado(nuevaCadena)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
