package metrics_test

import (
	"testing"

	"github.com/pchchv/bpdf/metrics"
	"github.com/stretchr/testify/assert"
)

func TestTime_Normalize(t *testing.T) {
	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		time := metrics.Time{
			Value: 3000,
			Scale: metrics.Nano,
		}
		ok := time.Normalize()

		assert.True(t, ok)
		assert.Equal(t, 3.0, time.Value)
		assert.Equal(t, metrics.Micro, time.Scale)
	})

	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Micro,
		}
		ok := time.Normalize()

		assert.True(t, ok)
		assert.Equal(t, 2.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})

	t.Run("when scale is milli, should return false", func(t *testing.T) {
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Milli,
		}
		ok := time.Normalize()

		assert.False(t, ok)
		assert.Equal(t, 2000.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})
}

func TestTime_String(t *testing.T) {
	time := metrics.Time{
		Value: 2000,
		Scale: metrics.Milli,
	}
	s := time.String()

	assert.Equal(t, "2000.00ms", s)
}

func TestSize_Normalize(t *testing.T) {
	t.Run("when scale is byte, should divide by 1000 and change to kilo", func(t *testing.T) {
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.Byte,
		}
		ok := size.Normalize()

		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.KiloByte, size.Scale)
	})

	t.Run("when scale is kilo, should divide by 1000 and change to mega", func(t *testing.T) {
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.KiloByte,
		}
		ok := size.Normalize()

		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.MegaByte, size.Scale)
	})

	t.Run("when scale is mega, should divide by 1000 and change to giga", func(t *testing.T) {
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.MegaByte,
		}
		ok := size.Normalize()

		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.GigaByte, size.Scale)
	})

	t.Run("when scale is giga, should return false", func(t *testing.T) {
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.GigaByte,
		}
		ok := size.Normalize()

		assert.False(t, ok)
		assert.Equal(t, 3000.0, size.Value)
		assert.Equal(t, metrics.GigaByte, size.Scale)
	})
}

func TestSize_String(t *testing.T) {
	size := metrics.Size{
		Value: 2000,
		Scale: metrics.KiloByte,
	}
	s := size.String()

	assert.Equal(t, "2000.00Kb", s)
}
