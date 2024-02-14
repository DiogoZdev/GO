package main

import (
	"fmt"
);

func main() {

	for x := 1; x <= 122; x++ {
		showNumber(x);
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