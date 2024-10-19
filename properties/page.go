package properties

const (
	Top         Place = "top"          // place in the top of the page
	LeftTop     Place = "left_top"     // place in the left top of the page
	RightTop    Place = "right_top"    // place in the right top of the page
	Bottom      Place = "bottom"       // place in the bottom of the page
	LeftBottom  Place = "left_bottom"  // place in the left bottom of the page
	RightBottom Place = "right_bottom" // place in the right bottom of the page
)

// Place is the representation of a place in a page.
type Place string

// IsValid checks if the place is valid.
func (p Place) IsValid() bool {
	return p == LeftTop || p == Top || p == RightTop ||
		p == LeftBottom || p == Bottom || p == RightBottom
}
