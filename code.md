package main

import (
	"fmt"
	"math"
)

// 1. An interface is just a set of method signatures.
//    Any type that has these methods automatically "is" a Shape.
//    There is no "implements" keyword in Go.

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Two unrelated structs.
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// 3. Give each struct the methods the interface asks for.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//  4. This function accepts ANY Shape. It doesn't know or care
//     whether it is a Rectangle, a Circle, or something added later.
func printInfo(s Shape) {
	fmt.Printf("Area: %.2f | Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 5. Type switch: find out the concrete type behind an interface
func describe(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("Rectangle %.1f x %.1f\n", v.Width, v.Height)
	case Circle:
		fmt.Printf("Circle with radius %.1f\n", v.Radius)
	default:
		fmt.Printf("Unknown shape")

	}
}

func main() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 7}

	// same function different concrete types
	printInfo(r)
	printInfo(c)

	// A slice of interface values can hold different types together

	shapes := []Shape{r, c, Rectangle{Width: 2, Height: 2}}

	total := 0.0

	for _, s := range shapes {
		describe(s)
		total += s.Area()
	}

	fmt.Printf("Total area: %.2f\n", total)

	// Type assertion: pull the concrete value back out.

	var s Shape = c
	if circle, ok := s.(Circle); ok {
		fmt.Println("It's a circle, radius:", circle.Radius)
	}

	// The empty interface (any) accepts a value of any type.
	var anything any = 42
	anything = "now a string"
	anything = r
	fmt.Println("anything holds:", anything)

}
