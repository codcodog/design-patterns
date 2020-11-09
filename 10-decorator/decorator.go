package decorator

import "fmt"

type API interface {
	call() string
}

type Base struct {
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) call() string {
	return "base call"
}

type Decorator struct {
	api API
}

func NewDecorator(api API) *Decorator {
	return &Decorator{api}
}

func (d *Decorator) call() string {
	return fmt.Sprintf("decorator %s", d.api.call())
}
