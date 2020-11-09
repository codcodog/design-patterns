package command

func ExampleLightOn() {
	b := NewButton(NewLightOn(&Light{}))
	b.Click()
	// Output: light on
}

func ExampleLightOff() {
	b := NewButton(NewLightOff(&Light{}))
	b.Click()
	// Output: light off
}

func ExampleTVOn() {
	b := NewButton(NewTVOn(&TV{}))
	b.Click()
	// Output: TV on
}

func ExampleTVOff() {
	b := NewButton(NewTVOff(&TV{}))
	b.Click()
	// Output: TV off
}
