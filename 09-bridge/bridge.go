package bridge

type Shape interface {
	shape() string
}

type DrawlAPI interface {
	drawl() string
}

type Circle struct {
	drawlMethod DrawlAPI
}

func NewCircle(method DrawlAPI) *Circle {
	return &Circle{
		drawlMethod: method,
	}
}

func (c *Circle) shape() string {
	return "circle"
}

type Rectangle struct {
	drawlMethod DrawlAPI
}

func NewRectangle(method DrawlAPI) *Rectangle {
	return &Rectangle{
		drawlMethod: method,
	}
}

func (r *Rectangle) shape() string {
	return "rectangle"
}

type Green struct {
}

func NewGreen() DrawlAPI {
	return &Green{}
}

func (g *Green) drawl() string {
	return "green"
}

type Red struct {
}

func NewRed() DrawlAPI {
	return &Red{}
}

func (r *Red) drawl() string {
	return "red"
}
