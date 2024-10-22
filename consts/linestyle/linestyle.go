// The linestyle package contains all line styles.
package linestyle

// Type is a representation of a line style style.
type LineStyle string

const (
	DefaultLineThickness float64   = 0.2 // default line style width in gofpdf
	Solid                LineStyle = "solid"
	Dashed               LineStyle = "dashed"
)
