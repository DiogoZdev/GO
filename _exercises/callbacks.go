package main

import "fmt"

func main() {
	c := math(10, square);
	fmt.Println(c)
}

func double (x int) int {
	return x+x
}

func half (x int) int {
	return x/2
}

func square (x int) int {
	return x*x
}

func math(val int, f func(g int) int) int {
	c := f(val);

	return c;
}