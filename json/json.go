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

	aa := Shape{
		Lados: 6,
		Tamanho: 2,
		Cores: []string{"red", "green", "blue"},
	}
	
	bb, _ := json.Marshal(aa)

	fmt.Println(bb)
}
