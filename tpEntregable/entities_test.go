package main

import (
	"entregableGoTudai2021/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChequeoDeCadena(t *testing.T) {
	cadena := model.NewCadenaEntrante("TX03ABC")
	tipo := model.GetTipo(cadena.Value)
	valor := model.GetValue(cadena.Value)
	largo := model.GetLength(cadena.Value)
	assert.Equal(t, model.ChequeoDeCadena(tipo, valor, largo), true, "La cadena es invalida")
}
