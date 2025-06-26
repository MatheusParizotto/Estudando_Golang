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

func printNome(nome string) string {
	return nome
}

func main() {
	fmt.Println(" O resultado da soma é:", soma(42, 13))
	fmt.Println(" O resultado da subtração é:", sub(42, 13))
	fmt.Println(" O resultado da multiplicação é:", mult(42, 13))

	fmt.Println(printNome("Steph"))
}
