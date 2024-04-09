package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {

	p := Person{
		name: "Kowalski",
		age: 18,
	}

	fmt.Println(p)

	change(&p)

	fmt.Print(p)

}

func change(p *Person) {
	fmt.Println("var address: ", &p)
	(*p).name = "Chaged Name"
}