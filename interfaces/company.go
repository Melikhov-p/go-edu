package interfaces

type Worker interface {
	Work(tasks []string)
	Settle(c *Company)
}

type Company struct {
	Name     string
	personal []Worker
}

func (c *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
	newbie.Settle(c)
}

func (c *Company) Personal() []Worker {
	return c.personal
}
