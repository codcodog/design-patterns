package facade

import "testing"

func TestFacade(t *testing.T) {
	computer := &Computer{}
	expect := "cpu start memory load hard drive read"
	if computer.Start() != expect {
		t.Fatalf("test facade failed. expect: %s, actual: %s", expect, computer.Start())
	}
}
