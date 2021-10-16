package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	teste := "123"
	fmt.Print(teste)
	if nome == "" {
		nome = "mundo"
	}

	if idioma == "espanhol" {
		return prefixoOlaEspanhol + nome
	}

	if idioma == "frances" {
		return prefixoOlaFrances + nome
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
