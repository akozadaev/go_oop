package person

import "fmt"

type Person struct {
	firstName string
	lastName  string
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
	return fmt.Sprintf("Привет, я %s %s", p.firstName, p.lastName)
}
