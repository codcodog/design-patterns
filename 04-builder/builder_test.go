package builder

import "testing"

func TestBuilder1(t *testing.T) {
	b := &Builder1{}
	d := NewDirector(b)
	product := d.BuildProduct()

	if product.a != "builder1 a" {
		t.Fatal("test builder1 failed.")
	}
	if product.b != "builder1 b" {
		t.Fatal("test builder1 failed.")
	}
	if product.c != "builder1 c" {
		t.Fatal("test builder1 failed.")
	}
}

func TestBuilder2(t *testing.T) {
	b := &Builder2{}
	d := NewDirector(b)
	product := d.BuildProduct()

	if product.a != "builder2 a" {
		t.Fatal("test builder2 failed.")
	}
	if product.b != "builder2 b" {
		t.Fatal("test builder2 failed.")
	}
	if product.c != "builder2 c" {
		t.Fatal("test builder2 failed.")
	}
}
