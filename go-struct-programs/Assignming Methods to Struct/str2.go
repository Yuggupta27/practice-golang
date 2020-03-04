package main

import "fmt"

func main() {
	rect1 := Rectangle{height: 10, width: 5}
	fmt.Println("Area of rectangle is =", rect1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect1 *Rectangle) area() float64 {

	return rect1.height * rect1.width

}
