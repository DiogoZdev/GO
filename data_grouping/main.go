package main

import (
	"fmt"
)

func main() {
	// x := []int{}
	// y := []int{0,1,2,3,4,5,6,7,8,9,10}

	// after_append := append(x, y...)
	// fmt.Println(after_append)

	// after_slice := after_append[2:8]
	// fmt.Println(after_slice)


	slice := []int{9,8,7,6,5,4,3,2,1,0}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}