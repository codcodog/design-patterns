package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	config1 := NewConfig()
	config2 := NewConfig()
	if config1 == nil || config2 == nil {
		t.Fatal("test singleton failed.")
	}
	if config1 != config2 {
		t.Fatal("test singleton failed.")
	}
}

func TestParalleSingleton(t *testing.T) {
	num := 1000
	configs := make([]*config, num)
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			configs[i] = NewConfig()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < num; i++ {
		if configs[i] != configs[i-1] {
			t.Fatal("test paralle singleton failed.")
		}
	}
}
