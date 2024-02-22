package main

import "fmt"

type Human interface {
    hello()
}

type Person struct {
    nome string
}

func (p Person) hello() {
    fmt.Println(p.nome + " says: Hello!")
}

type Worker struct {
    nome string
    age  int
}

func (a Worker) hello() {
    fmt.Println(a.nome + " says: Hello! I Work.")
}

func sayHello(h Human) {
    h.hello()
}

func main() {
    p1 := Person{nome: "John"}
    p2 := Worker{nome: "Doe", age: 18}

    sayHello(p1)
    sayHello(p2)
}
