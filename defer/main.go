package main

import "fmt"

// o defer statement deixa a execução por último na fila de execução da função

func main() {
	defer fmt.Println("1")

	fmt.Println("2")
	fmt.Println("3")
}