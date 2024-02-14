package main

import (
	"fmt"
)

func main() {
	x := []int{}
	y := []int{0,1,2,3,4,5,6,7,8,9,10}

	after_append := append(x, y...)
	fmt.Println(after_append)

	after_slice := after_append[2:8]
	fmt.Println(after_slice)
}