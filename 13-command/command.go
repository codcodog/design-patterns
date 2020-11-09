package command

import "fmt"

type Command interface {
	execute()
}

type Light struct {
}

func (l *Light) TurnOn() {
	fmt.Println("light on")
}

func (l *Light) TurnOff() {
	fmt.Println("light off")
}

// 开灯命令
type LightOn struct {
	*Light
}

func NewLightOn(light *Light) *LightOn {
	return &LightOn{light}
}

func (l *LightOn) execute() {
	l.TurnOn()
}

// 关灯命令
type LightOff struct {
	*Light
}

func NewLightOff(light *Light) *LightOff {
	return &LightOff{light}
}

func (l *LightOff) execute() {
	l.TurnOff()
}

type TV struct {
}

func (t *TV) TurnOn() {
	fmt.Println("TV on")
}

func (t *TV) TurnOff() {
	fmt.Println("TV off")
}

// 开电视命令
type TVOn struct {
	*TV
}

func NewTVOn(tv *TV) *TVOn {
	return &TVOn{tv}
}

func (t *TVOn) execute() {
	t.TurnOn()
}

// 关电视命令
type TVOff struct {
	*TV
}

func NewTVOff(tv *TV) *TVOff {
	return &TVOff{tv}
}

func (t *TVOff) execute() {
	t.TurnOff()
}

type Button struct {
	command Command
}

func NewButton(c Command) *Button {
	return &Button{c}
}

func (b *Button) Click() {
	b.command.execute()
}
