package abstractfactory

import "testing"

func TestWin(t *testing.T) {
	winFactory := &WinFactory{}
	button := winFactory.createButton()
	if button.createButton() != "win button" {
		t.Fatal("test win button failed.")
	}

	border := winFactory.createBorder()
	if border.createBorder() != "win border" {
		t.Fatal("test win border failed.")
	}
}

func TestMac(t *testing.T) {
	macFactory := &MacFactory{}
	button := macFactory.createButton()
	if button.createButton() != "mac button" {
		t.Fatal("test mac button failed.")
	}

	border := macFactory.createBorder()
	if border.createBorder() != "mac border" {
		t.Fatal("test mac border failed.")
	}
}
