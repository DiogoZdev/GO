package main

import (
	"fmt"
)

const (
	a = iota + 2025
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)	
}
