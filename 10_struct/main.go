package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	sobrenome string
	sorvetes []string
}

func main() {
	example1()
	example2()
	example3()
	example4()
}

type Person struct {
	name string
	age int
	license bool
}

type Equipment struct {
	name string
	value float64
}

type Machine struct {
	detail Equipment
	exists bool
}

func example1() {
	fmt.Println("------\n\n Struct simples e suas inicializações\n\n------")

	p1 := Person{
		age: 18,
		name: "Diogo",
		license: true,
	}

	p2 := Person{"Andressa", 18, true}

	fmt.Println(p1, p2)
}

func example2() {
	fmt.Println("------\n\n Nested Structs e suas inicializações\n\n------")

	data := Machine{
		detail: Equipment{
			name: "Laptop",
			value: 1000.00,
		},
		exists: true,
	}

	data2 := Machine{ Equipment{"Tablet", 500.00}, true}

	fmt.Println(data, data2)
	fmt.Println(data2.detail.value)
}

func example3() {
	fmt.Println("------\n\n Structs anonimos\n\n------")

	// anônimo pois não é declarado antes de ser inicializado e não há um tipo reutilizável

	player := struct {
		name string
		game string
	} {
		name: "Diogo",
		game: "GTA",
	}

	fmt.Println(player)
}

func example4 () {
	fmt.Println("------\n\n Structs, Maps and Nested Maps\n\n------")

	p1 := Pessoa{
		nome: "James",
		sobrenome: "Lara",
		sorvetes: []string{"Chocolate", "Morango"},
	}

	p2 := Pessoa{
		nome: "Croft",
		sobrenome: "Lara",
		sorvetes: []string{"Chocolate", "Morango"},
	}

	p3 := Pessoa{
		nome: "Pedepano",
		sobrenome: "Lara",
		sorvetes: []string{"ogro", "fumaça"},
	}

	mapa := map[string]Pessoa{
		p1.nome: p1,
		p2.nome: p2,
		p3.nome: p3,
	}

	display(mapa);
}

func display(data map[string]Pessoa) {

	for _, v := range data {
		fmt.Println(v.nome, v.sobrenome)
		for id, sor := range v.sorvetes {
			fmt.Println("\t", id+1, " - ", sor,)
		}
		fmt.Println()
	}
}
