package generator

import (
	"errors"
	"sync"
	"time"
)

// Engine is responsible for generating unique IDs using the provided IDGenerator while enforcing concurrency control.
type Engine struct {
	mu            sync.Mutex
	generator     *IDGenerator
	timeReference int64
}

// NewEngine creates and returns a new Engine instance using the provided IDGenerator for unique ID generation.
func NewEngine(generator *IDGenerator) *Engine {
	return &Engine{generator: generator}
}

// GetID generates a unique ID for the given time. It may reset the internal state when the timestamp changes.
func (e *Engine) GetID(t time.Time) (uint64, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.timeReference < t.UnixMilli() {
		e.generator.Reset()
		e.timeReference = t.UnixMilli()
	}

	id := e.generator.NextID(t)
	return id.GetValue()
}

// MustGetID generates a unique ID for the specified time or panics if no more IDs can be generated for the provided timestamp.
func (e *Engine) MustGetID(t time.Time) uint64 {
	for {
		id, err := e.GetID(t)
		var counterLimitErr CounterLimitError
		if errors.As(err, &counterLimitErr) {
			time.Sleep(time.Millisecond)
			continue
		}

		var timeLimitErr TimeLimitError
		if errors.As(err, &timeLimitErr) {
			panic("No more available bits for time")
		}

		return id
	}
}
