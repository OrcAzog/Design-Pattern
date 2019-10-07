package main

import "fmt"

type Hero interface {
	sayHello()
}

type Builder interface {
	build()
}

type Batman struct {
}

type GothamBuilder struct {
}

func (Batman) sayHello() {
	fmt.Print("I am Batman")
}

type City struct {
	hero Hero
	name string
}

func (c City) setName(name string) {
	c.name = name

}
func (c City) setHero(temp *Hero) {
	c.hero = *temp
}

func (GothamBuilder) build() City {
	return City{new(Batman), "Gotham"}
}
