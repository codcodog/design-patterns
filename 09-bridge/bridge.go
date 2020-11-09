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
