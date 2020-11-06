package adapter

// 鸟
type Bird interface {
	fly() string
	makeSound() string
}

// 麻雀
type Sparrow struct {
}

func (s *Sparrow) fly() string {
	return "sparrow fly"
}

func (s *Sparrow) makeSound() string {
	return "sparrow sound"
}

// 玩具鸭
type ToyDuck interface {
	squeak() string
}

// 塑料玩具鸭
type PlasticToyDuck struct {
}

func (p *PlasticToyDuck) squeak() string {
	return "toyduck squeak"
}

type BirdAdapter struct {
	bird Bird
}

func NewBirdAdapter(b Bird) *BirdAdapter {
	return &BirdAdapter{
		bird: b,
	}
}

func (b *BirdAdapter) squeak() string {
	return b.bird.makeSound()
}
