package main

import (
	"fmt"
	"math"
)

type test interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	r float64
}

func (rec Rectangle) area() float64 {
	return rec.height * rec.width
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	r1 := Rectangle{10, 20}
	c1 := Circle{5}
	fmt.Println(r1.area())
	fmt.Println(c1.area())
}
