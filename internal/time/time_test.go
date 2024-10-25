package time_test

import (
	"testing"
	buildtinTime "time"

	"github.com/pchchv/bpdf/internal/time"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeSpent(t *testing.T) {
	timeSpent := time.GetTimeSpent(func() {
		buildtinTime.Sleep(10 * buildtinTime.Millisecond)
	})

	assert.InDelta(t, float64(10*buildtinTime.Millisecond), timeSpent.Value, float64(2*buildtinTime.Millisecond))
}
