package main

import "fmt"

type Animal interface {
	Pet()
	Name() string
}

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) Pet() {
	fmt.Println("prrrrr")
}

type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Pet() {
	fmt.Println("woff woff")
}

func Compliment(a Animal) {
	fmt.Println("Good job", a.Name())
	a.Pet()
}

func main() {
	c := Cat{name: "Larry"}
	d := Dog{name: "Harry"}

	Compliment(c)
	Compliment(d)
}
