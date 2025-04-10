package main

type ICar interface {
	setName(name string)
	setPower(name int)
	getName() string
	getPower() int
}
