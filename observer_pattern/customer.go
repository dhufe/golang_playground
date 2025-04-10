package main

import (
	"fmt"
)

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending mail to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getId() string {
	return c.id
}
