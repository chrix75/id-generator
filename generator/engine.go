package generator

import (
	"errors"
	"sync"
	"time"
)

type Engine struct {
	mu            sync.Mutex
	generator     *IDGenerator
	timeReference int64
}

func NewEngine(generator *IDGenerator) *Engine {
	return &Engine{generator: generator}
}

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
