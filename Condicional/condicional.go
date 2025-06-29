package main

import (
	"fmt"
)

func main() {

	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero == 1 {
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor não é igual a 1 e nem 2")
	}

	if numero%2 == 0 {
		fmt.Println("O númeroo é par")
	} else {
		fmt.Println("O númeroo é impar")
	}

	// Operações
	// Soma: 1 + 1
	// Subtração: 2 - 1
	// Divisão: 10 / 2
	// Multiplicação: 2 * 2
	// Resto da divisão por x: 7%2 (resto da divisão por 2)

	/*
		fmt.Println(2 + 1)
		fmt.Println(2 - 1)
		fmt.Println(2 / 1)
		fmt.Println(2 * 2)
		fmt.Println(10 % 1)
	*/

}
