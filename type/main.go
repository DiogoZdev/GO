package main

import (
	"fmt"
	"reflect"
)

func main () {
	// usando o pkg reflect
	x := "a"
	fmt.Print("o tipo de ", x, " é ", reflect.TypeOf(x), "\n");


	// usando Sprintf
	y := "oi"
	typeOf := fmt.Sprintf("%T", y)

	fmt.Print("o tipo de ", y, " é ", typeOf, "\n");
}