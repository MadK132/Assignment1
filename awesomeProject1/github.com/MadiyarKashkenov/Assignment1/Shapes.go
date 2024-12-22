package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 2 * (s.Length + s.Length)
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	Per := t.Perimeter() / 2
	return math.Sqrt(Per * (Per - t.SideA) * (Per - t.SideB) * (Per - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func PrintShapeDetails(s Shape) {
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println("Area:", s.Area())
}
func main() {
	shapes := []Shape{
		Rectangle{Length: 3, Width: 4},
		Circle{Radius: 3},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
		Square{Length: 5},
	}
	for _, shape := range shapes {
		fmt.Println("Shape:", shape)
		PrintShapeDetails(shape)
	}
}
