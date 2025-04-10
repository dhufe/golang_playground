package main

type BMW struct {
	Car
}

func newBMW() ICar {
	return &BMW{
		Car: Car{
			name:  "BWM brand",
			power: 150,
		},
	}
}
