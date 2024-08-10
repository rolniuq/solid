package main

type GoodShape interface {
	Area() float64
}

type GoodCircle struct {
	Radius float64
}

func (g *GoodCircle) Area() float64 {
	return 3.14 * g.Radius * g.Radius
}

type GoodSquare struct {
	Side float64
}

func (g *GoodSquare) Area() float64 {
	return g.Side * g.Side
}

// -> with this we can add new shape: rectangle, we don't need to modify this code of Area

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func goodInit() {
	var shape GoodShape
	shapeWanted := "square"
	switch shapeWanted {
	case "circle":
		shape = &GoodCircle{
			Radius: 5,
		}
	case "rectangle":
		shape = &Rectangle{
			Width:  5,
			Height: 5,
		}
	case "square":
		shape = &GoodSquare{
			Side: 5,
		}
	}

	shape.Area()
}
