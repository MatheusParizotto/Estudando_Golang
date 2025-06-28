package main

import (
	"fmt"
)

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Steph", 28})
	fmt.Println(Pessoa{Nome: "Bento", Idade: 18})
	fmt.Println(Pessoa{"Joseph", 28})

	p1 := Pessoa{Nome: "Matheus", Idade: 23}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 24
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Gyo", Idade: 22}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// Struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
