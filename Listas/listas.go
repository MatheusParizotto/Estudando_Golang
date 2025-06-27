package main

import (
	"fmt"
)

func main() {

	var array [2]string
	array[0] = "Salve"
	array[1] = "Fellas!!"

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3])
	fmt.Println(numPrimos[:2])
	fmt.Println(numPrimos[1:])

	//var slice []string
	slice := make([]string, 5)

	slice[0] = "Hello"
	slice[1] = "World"

	fmt.Println(slice[0], slice[1])
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	slice[2] = "Aiai, esse tal de slice"
	fmt.Println(slice[2])

	fmt.Println(len(slice))

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16, 18)
	fmt.Println(numPares)

	// Listas

	// 1 - Arrays e Slices: Homogêneos
	// Todos os elementos tem o mesmo tipo
	// [1, 2, 3, 4, 5, 6] - []int
	// ["steph", "bento", "golang"] - []string

	// 2 - Maps: Heterogêneos
	// Pode misturar tipos
	// Estrutura chave - valir
	// [Key] = value
	// Chave tem um tipo, e o valor pode ter outro
	// map[string]int
	//	{ "steph": 28, "bento": 4 }
	// map[string]string
	//	{ "steph": "cardoso", "bento": "pereira" }

	// Array

	// Tamanho fixo, de zero ou mais elementos do mesmo tipo
	// Acessamos os valores com índice: a[0], a[1]...
	// Função embutida len() retorna o tamnaho do array
	// Por conta do tamanho fixo, não é tão usado. somente em casos específicos

	//Slice

	// Tipo o array, mas sem tamanho fixo
	// Acessamos os valores com índices: a[0], a[1]...
	// Função embutida len() retorna o tamnaho do slice
	// Função append() usada para adicionar valores

}
