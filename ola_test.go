package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Carsio")
	esperado := "Olá, Carsio"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
