package main

import (
	"fmt"
)

func main() {
	a := createClosure()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println()

	b := createClosure()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func createClosure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Na func closure, retornamos uma função que retorna um int. 
// Porém, o valor de x criado, está fora do escopo da função retonada,
// fazendo assim a criação de um novo escopo onde aquela variável existe apenas onde `closure()` é atribuído.

// Cada vez que uma nova variável recebe a função `closure()` 
// um novo tipo de `x := 0` é criado. 

// Assim, cada vez que uma variável recebe closure e sua chamada é feita,
// cada x existe em um escopo diferente.

// o `x` de `a` é um
// o `x` de `b` é outro