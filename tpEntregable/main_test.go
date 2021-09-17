package main //porque estoy en el mismo nivel

import "testing"
import "entregableGoTudai2021/model"

func TestMain(t *testing.T) {
	main()
}

func TestNewCadenaEntrante(t *testing.T) {
	model.NewCadenaEntrante("TX03ABC")
}

func TestgenerarResultado(t *testing.T) {
	generarResultado("TX03ABC")
}
