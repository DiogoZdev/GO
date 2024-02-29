package main

import "fmt"

type Shape interface {
	area() string
}

type Circle struct {
	radius float64
}

func (c Circle) area() string {
	PI := 3.1415
	return fmt.Sprintln("circle area: ", PI * c.radius * 2)
}

type Square struct {
	width float64
}

func (s Square) area() string {
	return fmt.Sprintln("square area: ", s.width * s.width)
}

func main() {
	c := Circle{radius: 2.5}
	s := Square{width: 5}

	shapeInfo(c)
	shapeInfo(s)
}

func shapeInfo(s Shape) {
	fmt.Println(s.area())
}