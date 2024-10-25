// Package metrics contains metrics models, constants and formatting.
package metrics

import "fmt"

const (
	Nano     TimeScale = "ns" // time scale in nanoseconds
	Micro    TimeScale = "μs" // time scale in microseconds
	Milli    TimeScale = "ms" // time scale in milliseconds
	Byte     SizeScale = "b"  // size scale in bytes
	KiloByte SizeScale = "Kb" // size scale in kilobytes
	MegaByte SizeScale = "Mb" // size scale in megabytes
	GigaByte SizeScale = "Gb" // size scale in gigabytes
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

// Normalize normalizes the size scale.
func (t *Size) Normalize() bool {
	if t.Scale == Byte {
		t.Scale = KiloByte
		t.Value /= 1000.0
		return true
	}

	if t.Scale == KiloByte {
		t.Scale = MegaByte
		t.Value /= 1000.0
		return true
	}

	if t.Scale == MegaByte {
		t.Scale = GigaByte
		t.Value /= 1000.0
		return true
	}

	return false
}

// String returns the size formatted.
func (t *Size) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Scale)
}

// TimeMetric is a time metric.
type TimeMetric struct {
	Key   string
	Times []*Time
	Avg   *Time
}

func (m *TimeMetric) hasGreaterThan1000(times []*Time) bool {
	for _, time := range times {
		if time.Value > 1000.0 {
			return true
		}
	}
	return false
}

// Normalize normalizes the time metric.
func (m *TimeMetric) Normalize() {
	greaterThan1000 := m.hasGreaterThan1000(m.Times)
	if greaterThan1000 {
		for _, time := range m.Times {
			done := time.Normalize()
			if !done {
				return
			}
		}
	}

	if greaterThan1000 {
		m.Normalize()
	}
}

// String returns the time metric formatted.
func (m *TimeMetric) String() (content string) {
	content += m.Key + " -> avg: " + m.Avg.String() + ", executions: ["
	for i, time := range m.Times {
		content += time.String()
		if i < len(m.Times)-1 {
			content += ", "
		}
	}

	content += "]"
	return content
}

// SizeMetric is a size metric.
type SizeMetric struct {
	Key  string
	Size Size
}

// Normalize normalizes the size metric.
func (s *SizeMetric) Normalize() {
	if s.Size.Value < 1000.0 {
		return
	}

	s.Size.Normalize()
	s.Normalize()
}

// String returns the size metric formatted.
func (s *SizeMetric) String() string {
	return s.Key + " -> " + s.Size.String()
}
