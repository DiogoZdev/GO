package main

import "fmt"

func main() {

	fmt.Println(fatorial(10))
}

func fatorial(n int64) int64 {
	if n == 0 {
		return 1
	}

	return n * fatorial(n-1)
}