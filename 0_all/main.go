package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}
type Duck struct{}

func (Dog) Speak() string  { return "Гав" }
func (Cat) Speak() string  { return "Мяу" }
func (Duck) Speak() string { return "Кря-кря" }

func main() {
	animals := []Animal{Dog{}, Cat{}, Duck{}}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
