package bridge

import (
	"fmt"
)

func ExampleGreenCircle() {
	c := NewCircle(NewGreen())
	fmt.Println(c.shape())
	fmt.Println(c.drawlMethod.drawl())
	// Output:
	// circle
	// green
}

func ExampleRedCircle() {
	c := NewCircle(NewRed())
	fmt.Println(c.shape())
	fmt.Println(c.drawlMethod.drawl())
	// Output:
	// circle
	// red
}

func ExampleRedRectangle() {
	r := NewRectangle(NewRed())
	fmt.Println(r.shape())
	fmt.Println(r.drawlMethod.drawl())
	// Output:
	// rectangle
	// red
}
