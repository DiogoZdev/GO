package main

import "fmt"

// Todos os valores ficãm armaenados em memória
// Toda localização de memória possiu endereço
// O ponteiro indica esse endereço.

func main() {

	x := 10
	// &<var> retorna o endereço em memória de <var>
	y := &x

	ShowPointerData(x)

	// tendo o endereço, é possível recuperar o valor com *<endereço>
	fmt.Println("o valor no endereço ", y, " é ", *y)
}

func ShowPointerData(x any) {
	fmt.Println("O endereço de x é: ", &x)
	fmt.Println("O valor de x é: ", x)
}