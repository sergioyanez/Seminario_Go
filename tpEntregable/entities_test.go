package main

import (
	"testing"

	"entregableGoTudai2021/model"
)

func TestChequeoDeCadena(t *testing.T) {
	cadena := model.NewCadenaEntrante("TX03ABC")
	tipo := model.GetTipo(cadena.Value)
	valor := model.GetValue(cadena.Value)
	largo := model.GetLength(cadena.Value)
	if !model.ChequeoDeCadena(tipo, valor, largo) {
		t.Error("La cadena es invalida")
	}
}
