package interfaces

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  *Company
}

func (p *Person) GrowUp() {
	p.Age++
}

func (p *Person) Work(tasks []string) {
	for _, t := range tasks {
		fmt.Println("Task done: ", t)
	}
}

func (p *Person) Settle(c *Company) {
	p.Job = c
}
