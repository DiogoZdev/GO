package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	addGoRoutines(100)
	// time.Sleep(time.Second*2)
}

func addGoRoutines(amount int) {
	wg.Add(amount)

	for i := 0; i < amount; i++ {
		n := i
		go func(x int) {
			fmt.Println("GoRoutine ", x+1)
		}(n)
		wg.Done()
	}

	wg.Wait()
}