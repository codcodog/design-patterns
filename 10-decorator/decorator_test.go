package decorator

import "fmt"

func ExampleDecorator() {
	d := NewDecorator(NewBase())
	fmt.Println(d.call())
	// Output: decorator base call
}
