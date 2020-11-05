package builder

type Builder interface {
	buildPartA()
	buildPartB()
	buildPartC()
	GetProduct() *Product
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) BuildProduct() *Product {
	d.builder.buildPartA()
	d.builder.buildPartB()
	d.builder.buildPartC()
	return d.builder.GetProduct()
}

type Product struct {
	a string
	b string
	c string
}

type Builder1 struct {
	a string
	b string
	c string
}

func (bb *Builder1) buildPartA() {
	bb.a = "builder1 a"
}
func (bb *Builder1) buildPartB() {
	bb.b = "builder1 b"
}
func (bb *Builder1) buildPartC() {
	bb.c = "builder1 c"
}
func (bb *Builder1) GetProduct() *Product {
	return &Product{
		a: bb.a,
		b: bb.b,
		c: bb.c,
	}
}

type Builder2 struct {
	a string
	b string
	c string
}

func (bb *Builder2) buildPartA() {
	bb.a = "builder2 a"
}
func (bb *Builder2) buildPartB() {
	bb.b = "builder2 b"
}
func (bb *Builder2) buildPartC() {
	bb.c = "builder2 c"
}
func (bb *Builder2) GetProduct() *Product {
	return &Product{
		a: bb.a,
		b: bb.b,
		c: bb.c,
	}
}
