package structs

import "math"

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle structure
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculate perimeter of Rectangle
func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area calculate area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle structure
type Circle struct {
	Radius float64
}

// Area calculate area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculate perimeter of Circle
func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}
