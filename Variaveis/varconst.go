package main

import (
	"fmt"
)

func main() {

	// Variáveis
	// var + nome da variável + tipo

	var nome string
	nome = "Bento"
	fmt.Println(nome)

	nome = "Steph"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "Cleber"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	// Constantes

	const idadeBento = 15
	fmt.Println(idadeBento)
}
