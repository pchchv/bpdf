// Package metrics contains metrics models, constants and formatting.
package metrics

type (
	TimeScale string
	SizeScale string
)

// Time scales.
type Time struct {
	Value float64
	Scale TimeScale
}
