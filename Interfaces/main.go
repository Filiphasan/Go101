package main

import "fmt"

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

func (r Rectangle) getArea() float64 {
	return r.length * r.width
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {

	var shapes []Shape
	square := Square{side: 5}
	rectangle := Rectangle{length: 10, width: 5}
	shapes = append(shapes, square)
	shapes = append(shapes, rectangle)
	for _, shape := range shapes {
		fmt.Println("Area:", shape.getArea())
		fmt.Println("Perimeter:", shape.getPerimeter())
	}

}
