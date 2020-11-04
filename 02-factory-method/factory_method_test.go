package factorymethod

import "testing"

func TestRoundButton(t *testing.T) {
	buttonFactory := &RoundButtonFactory{}
	button := buttonFactory.createButton()
	if button.shape() != "round button" {
		t.Fatal("test round button failed.")
	}
}

func TestRectangleButton(t *testing.T) {
	buttonFactory := &RectangleButtonFactory{}
	button := buttonFactory.createButton()
	if button.shape() != "rectangle button" {
		t.Fatal("test rectangle button failed.")
	}
}
