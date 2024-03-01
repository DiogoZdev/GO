package main

import "fmt"

func main() {
	c := createClosureCount()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	d := createClosureCount()
	fmt.Println(d())
}

func createClosureCount() func() int {

	x := 1

	return func() int {
		x++
		return x
	}
}