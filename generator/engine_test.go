package generator

import (
	"testing"
	"time"
)

func TestLotOfID(t *testing.T) {
	// given
	gen := NewGenerator(1)
	engine := NewEngine(gen)

	// when
	var previousValue uint64 = 0
	for range 100_000 {
		v := engine.MustGetID(time.Now())

		if v <= previousValue {
			t.Fatalf("%v should be greater than %d", v, previousValue)
		}
		previousValue = v
	}
}
