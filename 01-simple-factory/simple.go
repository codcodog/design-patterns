package simplefactory

type Button interface {
	shape() string
}

func NewButton(t string) Button {
	switch t {
	case "round":
		return &RoundButton{}
	case "rectangle":
		return &RectangleButton{}
	case "diamond":
		return &DiamondButton{}
	default:
		panic("error button type")
	}
}

type RoundButton struct {
}

func (r *RoundButton) shape() string {
	return "round button"
}

type RectangleButton struct {
}

func (r *RectangleButton) shape() string {
	return "rectangle button"
}

type DiamondButton struct {
}

func (d *DiamondButton) shape() string {
	return "diamond button"
}
