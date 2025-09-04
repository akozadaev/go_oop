package main

import (
	"fmt"

	"github.com/akozadaev/go_oop/2_method/person"
)

func main() {
	introduceString := person.Person{"Алексей", "Козадаев"}.Introduce()
	fmt.Println(introduceString)

	person1 := person.Person{FirstName: "Алексей", LastName: "Козадаев"}
	introduceString = person1.Introduce()
	fmt.Println(introduceString)
}
