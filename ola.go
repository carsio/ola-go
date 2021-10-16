package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		return prefixoOlaPortugues + "mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
