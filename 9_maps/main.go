package main

import "fmt"

func main() {
	x := map[string][]string{
		"Diogo": {
			"Programar",
			"Estudar",
			"Comer a batata",
		},
		"Andressa": {
			"Desenhar",
			"Dormir",
			"Escutar Heavy Metal",
		},
	}

	val, exists := x["Diogo"]

	fmt.Println(val)
	fmt.Println(exists)

	if exists {
		val2 := val[2]
		fmt.Println(val2)
	}
	
}