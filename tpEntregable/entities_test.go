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

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
		//	{"TX3", false, "", "", 3},
	}

	for _, testData := range cases {
		data, err := generarResultado(testData.Input)
		error2 := (err == nil)
		assert.Equal(t, data.Type, testData.Type)
		assert.Equal(t, data.Value, testData.Value)
		assert.Equal(t, data.Length, testData.Length)
		assert.Equal(t, error2, testData.Success)

	}
}
