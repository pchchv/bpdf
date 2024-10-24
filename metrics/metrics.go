// Package metrics contains metrics models, constants and formatting.
package metrics

import "fmt"

const (
	Nano  TimeScale = "ns" // time scale in nanoseconds
	Micro TimeScale = "μs" // time scale in microseconds
	Milli TimeScale = "ms" // time scale in milliseconds
)

type (
	TimeScale string
	SizeScale string
)

// Time scales.
type Time struct {
	Value float64
	Scale TimeScale
}

// Normalize normalizes the time scale.
func (t *Time) Normalize() bool {
	if t.Scale == Nano {
		t.Scale = Micro
		t.Value /= 1000.0
		return true
	}

	if t.Scale == Micro {
		t.Scale = Milli
		t.Value /= 1000.0
		return true
	}

	return false
}

// String returns the time formatted.
func (t *Time) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Scale)
}

// Size scales.
type Size struct {
	Value float64
	Scale SizeScale
}
