package main

import (
	"fmt"
)

func getCar(carBrand string) (ICar, error) {
	if carBrand == "ford" {
		return newFord(), nil
	}
	if carBrand == "bmw" {
		return newBMW(), nil
	}

	return nil, fmt.Errorf("Wrong car brand passed")
}
