package interfaces

import (
	"fmt"
)

func main() {
	p1 := &Person{
		Name: "John Doe",
		Age:  24,
	}

	c1 := Company{Name: "BigTech"}

	fmt.Println(p1.Job)

	c1.Hire(p1)

	fmt.Println(p1.Job)
}
