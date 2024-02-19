package main

import "fmt"

type shape interface {
	getArea() float32
	getPerimeter() float32
}

type triangle struct {
	base   float32
	height float32
}

func (t triangle) getArea() float32 {
	return t.base * t.height / 2.0
}

func (t triangle) getPerimeter() float32 {
	return 0.0
}

type rectangle struct {
	width  float32
	height float32
}

func (r rectangle) getArea() float32 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float32 {
	return (r.width + r.height) * 2.0
}

func showShapeProperties(s shape) {
	fmt.Printf("\nArea: %f", s.getArea())
	fmt.Printf("\nPerimeter: %f", s.getPerimeter())
}

func main() {
	myTri := triangle{3.0, 3.0}
	myRec := rectangle{4.0, 2.0}
	showShapeProperties(myTri)
	showShapeProperties(myRec)
}
