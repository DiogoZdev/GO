package main

import "fmt"

// Em GO não existem classes, tudo são funções.
// Caso eu queira especificar que uma struct Pessoa tenha um método de saudação por exemplo
// devo associar o método de saudação a uma struct.

type Pessoa struct {
	nome string
	idade int
}

func (p Pessoa) saudacao() {
	fmt.Println(p.nome, "diz: Olá!")
}

func main() {

	p := Pessoa{
		nome: "Diogo",
		idade: 18,
	}

	p.saudacao()
}