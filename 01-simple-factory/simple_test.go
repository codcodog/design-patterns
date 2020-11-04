package simplefactory

import "testing"

func TestRoundButton(t *testing.T) {
	button := NewButton("round")
	if button.shape() != "round button" {
		t.Fatal("round button test failed.")
	}
}

func TestRectangleButton(t *testing.T) {
	button := NewButton("rectangle")
	if button.shape() != "rectangle button" {
		t.Fatal("rectangle button test failed.")
	}
}

func TestDiamondButton(t *testing.T) {
	button := NewButton("diamond")
	if button.shape() != "diamond button" {
		t.Fatal("diamond button test failed.")
	}
}
