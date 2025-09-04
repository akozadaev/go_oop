package main

import (
	person1 "compose/person"
	"fmt"
)

func main() {
	person := person1.Person{}
	person.SetFirstName("Алексей")
	person.SetLastName("Козадаев")
	person.SetAddress(person1.Address{"Тамбов"})
	fmt.Println(person.Introduce())
	fmt.Println("=== Черезе геттеры ===")
	fmt.Println(person.GetFirstName())
	fmt.Println(person.GetLastName())

}
