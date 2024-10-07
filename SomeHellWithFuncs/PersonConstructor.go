package SomeHellWithFuncs

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Age       int
	Birthdate time.Time
}

type option func(*Person)

func NewPerson(options ...option) *Person {
	var p *Person = &Person{
		Name:      "John",
		Age:       18,
		Birthdate: time.Now(),
	}

	for _, m := range options {
		m(p)
	}

	return p
}

func name(name string) option {
	return func(p *Person) {
		p.Name = name
	}
}

func age(age int) option {
	return func(p *Person) {
		p.Age = age
	}
}

func birthdate(birthdate time.Time) option {
	return func(p *Person) {
		p.Birthdate = birthdate
	}
}

func main() {
	var p1 *Person = NewPerson()
	var p2 *Person = NewPerson(name("Paul"), age(8))
	var p3 *Person = NewPerson(name("Morty"), age(88), birthdate(time.Now()))

	fmt.Println(p1, p2, p3)
}
