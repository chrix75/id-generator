package generator

import "fmt"

const timeLimitValue uint64 = 0x0000040000000000
const counterLimitValue uint16 = 0x03FFF

type CounterLimitError struct {
	currentValue uint16
}

func (e CounterLimitError) Error() string {
	return fmt.Sprintf("failed to generate an ID, counter limit reached (limit: %d, current: %d)", counterLimitValue, e.currentValue)
}

type TimeLimitError struct {
	currentTime uint64
}

func (e TimeLimitError) Error() string {
	return fmt.Sprintf("failed to generate an ID, time limit reached (limit: %d, current: %d)", timeLimitValue, e.currentTime)
}

func acceptTimeValue(timeValue uint64) bool {
	return timeValue < timeLimitValue
}

func acceptCounterValue(counter uint16) bool {
	return counter < counterLimitValue
}
