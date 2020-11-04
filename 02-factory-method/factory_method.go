package factorymethod

type ButtonFactory interface {
	createButton() Button
}

type Button interface {
	shape() string
}

type RoundButtonFactory struct {
}

func (r *RoundButtonFactory) createButton() Button {
	return &RoundButton{}
}

type RoundButton struct {
}

func (r *RoundButton) shape() string {
	return "round button"
}

type RectangleButtonFactory struct {
}

func (r *RectangleButtonFactory) createButton() Button {
	return &RectangleButton{}
}

type RectangleButton struct {
}

func (r *RectangleButton) shape() string {
	return "rectangle button"
}
