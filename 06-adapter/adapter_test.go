package adapter

import "testing"

func TestAdapter(t *testing.T) {
	sparrow := &Sparrow{}
	expect := "sparrow sound"

	birdAdapter := NewBirdAdapter(sparrow)
	if birdAdapter.squeak() != expect {
		t.Fatalf("test adapter failed. expect: %s, actual: %s", expect, birdAdapter.squeak())
	}
}
