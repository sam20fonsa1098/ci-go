package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 10)
	}
}
