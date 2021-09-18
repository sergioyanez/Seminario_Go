package main

import (
	"entregableGoTudai2021/model"
	"errors"
	"fmt"
)

func generarResultado(cad string) (model.Result, error) {
	cadena := model.NewCadenaEntrante(cad)
	tipo := model.GetTipo(cadena.Value)
	valor := model.GetValue(cadena.Value)
	largo := model.GetLength(cadena.Value)

	if model.ChequeoDeCadena(tipo, valor, largo) {

		return model.NewResult(tipo, largo, valor), nil

	}
	return model.Result{}, errors.New("cadena Invalida")

}

func main() {
	cadena := model.NewCadenaEntrante("NN03487")

	fmt.Println(model.GetTipo(cadena.Value))
	fmt.Println(model.GetValue(cadena.Value))
	fmt.Println(model.GetLength(cadena.Value))
	fmt.Println(generarResultado("NN03487"))

}
