package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (sq square) area() float64 {
	return sq.side * sq.side
}

func (cl circle) area() float64 {
	return math.Pi * cl.radius * cl.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{radius: 10}
	s := square{side: 25}

	info(c)
	info(s)

}
