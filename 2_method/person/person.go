package person

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// TODO p *Person - не забыть показать!!
func (p Person) Introduce() string {
	return fmt.Sprintf("Привет, я %s %s", p.FirstName, p.LastName)
}
