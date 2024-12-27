package generator

import (
	"time"
)

var epoch = time.Date(2020, 6, 23, 0, 0, 0, 0, time.UTC)

const timeValueMask uint64 = 0x000003FFFFFFFFFF

type IDGenerator struct {
	machineID uint8
	counter   uint16
}

type ID struct {
	MachineID uint8
	TimeValue uint64
	Counter   uint16
}

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

func (g *IDGenerator) NextID(currentTime time.Time) ID {
	g.counter++
	return ID{
		MachineID: g.machineID,
		TimeValue: uint64(currentTime.Sub(epoch).Milliseconds()),
		Counter:   g.counter,
	}
}

func (g *IDGenerator) Reset() {
	g.counter = 0
}

func NewGenerator(machineID uint8) *IDGenerator {
	return &IDGenerator{
		machineID: machineID,
	}
}
