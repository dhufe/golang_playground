package main

type Car struct {
	name  string
	power int
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) setPower(power int) {
	c.power = power
}

func (c *Car) getPower() int {
	return c.power
}
