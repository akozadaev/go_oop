package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	A float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return s.A * s.A
}

func main() {
	figure := Shape(&Circle{10})
	figure2 := Shape(&Square{10})
	fmt.Println(figure.Area())
	fmt.Println(figure2.Area())
}
