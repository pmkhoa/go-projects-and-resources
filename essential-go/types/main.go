package main

import "fmt"

type point struct {
	x, y int
}

func NewPoint(x, y int) *point {
	return &point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	p := point{20, 40}
	fmt.Printf("p has x: %d and y: %d\n", p.x, p.y)

	p2 := point{x: 20, y: 50}
	fmt.Printf("p2 has x: %d and y: %d\n", p2.x, p2.y)

	r := rect{
		pos:    p2,
		width:  8,
		height: 12,
	}
	fmt.Println("rect: ", r)

	area := r.area()
	fmt.Printf("area: %d\n", area)

}
