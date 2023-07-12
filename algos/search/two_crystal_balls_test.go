package search

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(1000)
	data := make([]bool, 1000)
	for i := 0; i < 1000; i++ {
		if i < idx {
			data[i] = false
		} else {
			data[i] = true
		}
	}

	result := TwoCrystalBalls(data)
	if result != idx {
		t.Errorf("TwoCrystalBalls(): expected %v, actual %v", idx, result)
	}

	data = make([]bool, 100)
	result = TwoCrystalBalls(data)
	if result != -1 {
		t.Errorf("TwoCrystalBalls(): expected %v, actual %v", -1, result)
	}

	data = make([]bool, 100)
	for i := 0; i < 100; i++ {
		data[i] = true
	}
	result = TwoCrystalBalls(data)
	if result != 0 {
		t.Errorf("TwoCrystalBalls(): expected %v, actual %v", 0, result)
	}
}
