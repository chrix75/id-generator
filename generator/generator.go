package generator

import (
	"time"
)

// epoch represents the reference time used as the starting point for calculating elapsed time in ID generation.
var epoch = time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)

const timeValueMask uint64 = 0x000003FFFFFFFFFF

// IDGenerator is a type responsible for generating unique IDs based on a machine ID, timestamp, and counter.
// It uses a machine ID to ensure uniqueness across multiple instances and a counter for sequential ID generation.
// The counter is resettable, allowing IDs to restart from the beginning while maintaining uniqueness within the scope.
type IDGenerator struct {
	machineID uint8
	counter   uint16
}

// ID represents a unique identifier composed of a machine ID, a timestamp, and a counter value.
type ID struct {
	MachineID uint8
	TimeValue uint64
	Counter   uint16
}

// GetValue computes and returns the unique identifier as a uint64 value, or an error if validation fails.
func (i ID) GetValue() (uint64, error) {
	if err := i.validateData(); err != nil {
		return 0, err
	}

	var value uint64
	value = i.TimeValue & timeValueMask

	value = value << 22

	var machineID = uint64(i.MachineID) << 21
	value |= machineID

	var counter = uint64(i.Counter)
	value |= counter

	return value, nil
}

func (i ID) validateData() error {
	if !acceptTimeValue(i.TimeValue) {
		return TimeLimitError{i.TimeValue}
	}

	if !acceptCounterValue(i.Counter) {
		return CounterLimitError{i.Counter}
	}

	return nil
}

// NextID generates a new unique ID using the current machine ID, time since the epoch, and an incrementing counter.
func (g *IDGenerator) NextID(currentTime time.Time) ID {
	g.counter++
	return ID{
		MachineID: g.machineID,
		TimeValue: uint64(currentTime.Sub(epoch).Milliseconds()),
		Counter:   g.counter,
	}
}

// Reset sets the counter to 0, allowing the ID generation to restart sequentially.
func (g *IDGenerator) Reset() {
	g.counter = 0
}

// NewGenerator initializes and returns a pointer to an IDGenerator with the specified machine ID.
func NewGenerator(machineID uint8) *IDGenerator {
	return &IDGenerator{
		machineID: machineID,
	}
}
