package observer

type Subject struct {
	observers []Observer
	state     string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
		state:     "init",
	}
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.update(s.state)
	}
}

func (s *Subject) UpdateState(state string) {
	s.state = state
	s.notify()
}

type Observer interface {
	update(string)
}

type Observer1 struct {
	state string
}

func NewObserver1() *Observer1 {
	return &Observer1{"init"}
}

func (o *Observer1) update(state string) {
	o.state = state
}

type Observer2 struct {
	state string
}

func NewObserver2() *Observer2 {
	return &Observer2{"init"}
}

func (o *Observer2) update(state string) {
	o.state = state
}

type Observer3 struct {
	state string
}

func NewObserver3() *Observer3 {
	return &Observer3{"init"}
}

func (o *Observer3) update(state string) {
	o.state = state
}
