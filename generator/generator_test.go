package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateUniqueID(t *testing.T) {
	// given
	gen := NewGenerator(1)
	currentTime := time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)

	// when
	id := gen.NextID(currentTime)

	// then
	expectedID := ID{
		MachineID: 1,
		TimeValue: 1728000000,
		Counter:   1,
	}

	assert.Equal(t, expectedID, id)
}

func TestGenerateTwoUniquesID(t *testing.T) {
	// given
	gen := NewGenerator(1)
	currentTime := time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)

	// when
	_ = gen.NextID(currentTime)
	id := gen.NextID(currentTime)

	// then
	expectedID := ID{
		MachineID: 1,
		TimeValue: 1728000000,
		Counter:   2,
	}

	assert.Equal(t, expectedID, id)
}

func TestResetCounter(t *testing.T) {
	// given
	gen := NewGenerator(1)
	currentTime := time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)

	// when
	_ = gen.NextID(currentTime)
	_ = gen.NextID(currentTime)
	gen.Reset()
	id := gen.NextID(currentTime)

	// then
	expectedID := ID{
		MachineID: 1,
		TimeValue: 1728000000,
		Counter:   1,
	}

	assert.Equal(t, expectedID, id)
}

func TestBuildIDValue(t *testing.T) {
	// given
	id := ID{
		MachineID: 1,
		TimeValue: 1728000000,
		Counter:   1,
	}

	// when
	value, _ := id.GetValue()

	// then
	expectedValue := uint64(7247757314097153)
	assert.Equal(t, expectedValue, value)
}

func TestSameCounterOnDifferentMachines(t *testing.T) {
	// given
	id1 := ID{
		MachineID: 1,
		TimeValue: 141868800000,
		Counter:   1,
	}

	id2 := ID{
		MachineID: 2,
		TimeValue: 141868800000,
		Counter:   1,
	}

	// when
	value1, _ := id1.GetValue()
	value2, _ := id2.GetValue()

	// then
	assert.NotEqual(t, value2, value1)
}

func TestOrderedID(t *testing.T) {
	tests := []struct {
		name      string
		timeValue uint64
		counter   uint16
	}{
		{"same time but inc counter", 141868800000, 2},
		{"same counter but inc time", 141868800001, 1},
	}

	refID := ID{
		MachineID: 1,
		TimeValue: 141868800000,
		Counter:   1,
	}

	refValue, _ := refID.GetValue()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id := ID{
				MachineID: 1,
				TimeValue: test.timeValue,
				Counter:   test.counter,
			}

			value, _ := id.GetValue()

			assert.Greater(t, value, refValue)
		})
	}
}

func TestCheckTimeLimit(t *testing.T) {
	tests := []struct {
		name          string
		timeValue     uint64
		expectedError bool
	}{
		{"time limit reached", 4398046511104, true},
		{"time limit not reached", 4398046511103, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id := ID{
				MachineID: 1,
				TimeValue: test.timeValue,
				Counter:   1,
			}
			_, err := id.GetValue()

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckCounterLimit(t *testing.T) {
	tests := []struct {
		name          string
		counter       uint16
		expectedError bool
	}{
		{"counter limit reached", 16383, true},
		{"counter limit not reached", 16382, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id := ID{
				MachineID: 1,
				TimeValue: 1,
				Counter:   test.counter,
			}

			_, err := id.GetValue()

			if test.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func BenchmarkGetValue(b *testing.B) {
	id := ID{
		MachineID: 1,
		TimeValue: 141868800000,
		Counter:   1,
	}

	for i := 0; i < b.N; i++ {
		_, _ = id.GetValue()
	}
}
