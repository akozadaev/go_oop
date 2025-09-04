package main

import (
	"fmt"
	person1 "incaps/person"
)

func main() {
	person := person1.Person{}
	person.SetFirstName("Алексей")
	person.SetLastName("Козадаев")
	fmt.Println(person.Introduce())
	fmt.Println("=== Черезе геттеры ===")
	fmt.Println(person.GetFirstName())
	fmt.Println(person.GetLastName())

}
