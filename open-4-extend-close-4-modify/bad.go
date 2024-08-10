package main

// example we have shapes and we want to calculate area of them

type ShapeType int

const (
	Circle ShapeType = iota
	Square
)

type Shape struct {
	ShapeType ShapeType
	Radius    float64
	Side      float64
}

func (s *Shape) Area() float64 {
	switch s.ShapeType {
	case Circle:
		return 3.14 * s.Radius * s.Radius
	case Square:
		return s.Side * s.Side
	}
	return 0
}

// so for now if we want to add new shape: rectangle, we need to modify this code of Area -> suck
// -> use OCP

func badInit() {
	c := Shape{
		ShapeType: Circle,
		Radius:    5,
	}
	c.Area()

	s := Shape{
		ShapeType: Square,
		Side:      5,
	}
	s.Area()
}
