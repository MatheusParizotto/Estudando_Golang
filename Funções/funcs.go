package main

import (
	"fmt"
)

func soma(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

// Função começando com letra minúscula
// Função PRIVADA
// Só pode ser utilizada dentro do próprio pacote

func printNome(nome string) string {
	return nome
}

// Função começando com letra maiúscul
// Função PÚBLICA
// Pode ser utilizada em outros pacotes

func PrintNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

func main() {
	fmt.Println(" O resultado da soma é:", soma(42, 13))
	fmt.Println(" O resultado da subtração é:", sub(42, 13))
	fmt.Println(" O resultado da multiplicação é:", mult(42, 13))

	fmt.Println(printNome("Steph"))

	nome, _ := PrintNomeCompleto("Steph", "Silva")

	fmt.Println("Nome:", nome)
}
