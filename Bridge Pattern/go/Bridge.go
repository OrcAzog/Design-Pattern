package awesomeProject2

import "fmt"

type Hero interface {
	sayHello()
}

type Batman struct {
}

type GothamBuilder struct {
}

func (Batman) sayHello() {
	fmt.Print("I am Batman")
}

type City struct {
	name string
}

func (c City) heroSpeech(hero Hero) {
	fmt.Print("这里是", c.name)
	hero.sayHello()
}
func (c City) setName(name string) {
	c.name = name

}
