package main

import (
	"fmt"
	"math"
)

type perimeter interface {
	peri() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) peri() float64 {
	return r.width * r.height
}

func (c Circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rec := Rectangle{10, 20}
	cir := Circle{5}
	fmt.Println(rec.peri())
	fmt.Println(cir.peri())
}
