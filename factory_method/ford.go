package main

type Ford struct {
	Car
}

func newFord() ICar {
	return &Ford{
		Car: Car{
			name:  "Ford brand",
			power: 500,
		},
	}
}
