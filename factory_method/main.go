package main

import (
	"fmt"
)

func main() {
	ford, _ := getCar("ford")
	bmw, _ := getCar("bmw")

	printDetails(bmw)
	printDetails(ford)
}

func printDetails(c ICar) {
	fmt.Printf("Car brand: %s", c.getName())
	fmt.Println()
	fmt.Printf("Horse power: %d PS", c.getPower())
	fmt.Println()
}
