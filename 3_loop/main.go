package main

import (
	"fmt"
);

func main() {

	// for simples
	for x := 1; x <= 122; x++ {
		showNumber(x);
	}

	// for com chave valor usando range
	list := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	for k, v := range list{
		fmt.Println(k, v)
	}
}

func showNumber(n int) {
	switch {
		case n < 18:
			fmt.Printf("%d menor que 18 \n", n)
		case n >= 18 && n < 65:
			fmt.Printf("%d maior que 18 \n", n)
		default:
			fmt.Printf("%d maior que 65 \n", n)
	}
}