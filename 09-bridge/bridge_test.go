package bridge

import "testing"

func TestBridge(t *testing.T) {
	c1 := NewCircle(NewGreen())
	if c1.shape() != "circle" {
		t.Fatal("test shape failed.")
	}
	if c1.drawlMethod.drawl() != "green" {
		t.Fatal("test drawl failed.")
	}

	c2 := NewCircle(NewRed())
	if c2.shape() != "circle" {
		t.Fatal("test shape failed.")
	}
	if c2.drawlMethod.drawl() != "red" {
		t.Fatal("test drawl failed.")
	}
}
