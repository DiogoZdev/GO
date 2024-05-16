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
	fmt.Println("\n================================== string")
	v1 := "oi"
	type1 := fmt.Sprintf("%T", v1)
	fmt.Println("o tipo de ", v1, " é ", type1);

	fmt.Println("\n================================== int")
	v2 := 21
	type2 := fmt.Sprintf("%T", v2)
	fmt.Println("o tipo de ", v2, " é ", type2);

	fmt.Println("\n================================== float")
	v3 := 22.88
	type3 := fmt.Sprintf("%T", v3)
	fmt.Println("o tipo de ", v3, " é ", type3);

	fmt.Println("\n================================== bool")
	v4 := false
	type4 := fmt.Sprintf("%T", v4)
	fmt.Println("o tipo de ", v4, " é ", type4);

	fmt.Println("\n================================== byte slice")
	v5 := []byte("asdasd")
	type5 := fmt.Sprintf("%T", v5)
	fmt.Println("o tipo de ", v5, " é ", type5);
}