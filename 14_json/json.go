package main

import (
	"encoding/json"
	"fmt"
)

type Shape struct {
	Lados   int      `json:"Sides"`
	Tamanho int      `json:"Size"`
	Cores   []string `json:"Colors"`
}

func main() {

	shape1 := Shape{
		Lados: 6,
		Tamanho: 2,
		Cores: []string{"red", "green", "blue"},
	}
	
	// json.Marshal - retorna um array de bits, deve ser convertido para string
	jsonString, _ := json.Marshal(shape1)

	fmt.Println(string(jsonString))


	// Unmarshal - transforma o json em uma struct, podendo apenas retornar erro.
	var shape2 Shape
	// aqui apontamos o valor de destino no endereço da memória de shape2
	err := json.Unmarshal(jsonString, &shape2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("jsonStruct", shape2)
	fmt.Printf("jsonStruct address: %p \n", &shape2)
}
