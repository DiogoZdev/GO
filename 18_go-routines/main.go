package main

import (
	"fmt"
)

func main() {
	go countOddToHundred()
	countPairToHundred()
}

func countPairToHundred() {
	for i := 0; i < 1000; i++ {
		if (i % 2 == 0) {
			fmt.Println("FUNC 1", i)
		}
	}
}

func countOddToHundred() {
	for i := 0; i < 1000; i++ {
		if (i % 2 != 0) {
			fmt.Println("FUNC 2",i)
		}
	}
}