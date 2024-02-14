package main

import (
	"fmt"
)

func main() {
	x := 10

	showValues(x)

	y := x << 1

	showValues(y)
}

func showValues(n int) {
	fmt.Printf("decimal: %d, binary: %b, hex: %#x \n\n", n, n, n)
}
