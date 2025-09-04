package person

import "fmt"

type Address struct {
	City string
}

type Person struct {
	firstName string
	lastName  string
	address   Address
}

func (p *Person) GetAddress() Address {
	return p.address
}

func (p *Person) SetAddress(address Address) {
	p.address = address
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}

func (p *Person) SetLastName(name string) {
	p.lastName = name
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("Привет, я %s %s из %sа", p.firstName, p.lastName, p.address.City)
}
