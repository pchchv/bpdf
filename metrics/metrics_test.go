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
