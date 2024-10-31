// Package pagesize contains all default page sizes.
package pagesize

const (
	A1                  Size = "a1"
	A2                  Size = "a2"
	A3                  Size = "a3"
	A4                  Size = "a4"
	A5                  Size = "a5"
	A6                  Size = "a6"
	Letter              Size = "letter"
	Legal               Size = "legal"
	Tabloid             Size = "tabloid"
	DefaultTopMargin         = 10.0
	DefaultLeftMargin        = 10.0
	DefaultRightMargin       = 10.0
	DefaultBottomMargin      = 20.0025
	MinTopMargin             = 0.0
	MinLeftMargin            = 0.0
	MinRightMargin           = 0.0
	MinBottomMargin          = 0.0
	DefaultFontSize          = 10.0
	DefaultMaxGridSum        = 12.0
)

// Size is a representation of a page size.
type Size string

// GetDimensions returns the width and height of the page size.
func GetDimensions(pageSize Size) (float64, float64) {
	switch pageSize {
	case A1:
		return 594.0, 841.0
	case A2:
		return 419.9, 594.0
	case A3:
		return 297.0, 419.9
	case A5:
		return 148.4, 210.0
	case A6:
		return 105.0, 148.5
	case Letter:
		return 215.9, 279.4
	case Legal:
		return 215.9, 355.6
	case Tabloid:
		return 279.4, 431.8
	default: // A4
		return 210.0, 297.0
	}
}
