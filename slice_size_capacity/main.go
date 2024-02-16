package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 15)
	fmt.Println(x)

	x = []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x)

	x = append(x, 52, 53)
		
	fmt.Println("length: ", len(x), " | capacity: ", cap(x))
	fmt.Println(x)

	fmt.Println(x[2:6])
}
