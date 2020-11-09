package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	s := NewSubject()
	ob1 := NewObserver1()
	ob2 := NewObserver2()
	ob3 := NewObserver3()
	if ob1.state != "init" ||
		ob2.state != "init" ||
		ob3.state != "init" {
		t.Fatal("test observer failed.")
	}

	s.Attach(ob1)
	s.Attach(ob2)
	s.Attach(ob3)
	s.UpdateState("ing")

	if ob1.state != "ing" ||
		ob2.state != "ing" ||
		ob3.state != "ing" {
		t.Fatal("test observer failed.")
	}
}
