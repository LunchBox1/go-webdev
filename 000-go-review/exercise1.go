package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	area := c.radius * c.radius
	return area
}

func (s square) area() float64 {
	area := s.length * s.width
	return area
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	c := circle{3.0}
	s := square{3.0, 2.0}

	info(c)
	info(s)
}
