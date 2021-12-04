package main

import "testing"
import "./math"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30);
	}
}