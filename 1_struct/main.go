package main

import "fmt"

type Animal struct {
	Name    string
	Species string
	Age     int
}

func main() {
	animal := Animal{"Дик", "Собака", 5}
	animalCat := Animal{"Майя", "Кошка", 1}
	fmt.Println(animal)
	fmt.Println(
		fmt.Sprintf("Кличка: %s,\n что за зверь: %s,\n возраст: %d год (лет)", animalCat.Name, animalCat.Species, animalCat.Age),
	)
}
