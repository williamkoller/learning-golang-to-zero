package main

import (
	"fmt"
	"math"
)

type recipe interface {
	area() float64
}

func writeRadius(r recipe) {
	fmt.Printf("The area of recipe is %0.2f\n", r.area())
}

type rectangle struct {
	height float64
	width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}


func main() {
	r := rectangle{10, 15}
	writeRadius(r)

	c := circle{10}

	writeRadius(c)
}