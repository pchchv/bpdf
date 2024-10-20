package properties

import (
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/breakline"
	"github.com/pchchv/bpdf/consts/fontstyle"
)

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

// PageNumber have attributes of page number
type PageNumber struct {
	// Pattern is the string pattern which will be used to apply the page count component.
	Pattern string
	// Place defines where the page count component will be placed.
	Place Place
	// Family defines which font family will be applied to page count.
	Family string
	// Style defines which font style will be applied to page count.
	Style fontstyle.Fontstyle
	// Size defines which font size will be applied to page count.
	Size float64
	// Color defines which will be applied to page count.
	Color *Color
}

// GetNumberTextProp returns the Text properties of the page number.
func (p *PageNumber) GetNumberTextProp(height float64) *Text {
	text := &Text{
		Family: p.Family,
		Style:  p.Style,
		Size:   p.Size,
		Color:  p.Color,
		Align:  align.Center,
	}

	if p.Place == LeftBottom || p.Place == LeftTop {
		text.Align = align.Left
	} else if p.Place == RightBottom || p.Place == RightTop {
		text.Align = align.Right
	}

	if p.Place == RightBottom || p.Place == Bottom || p.Place == LeftBottom {
		text.Top = height
	}

	text.BreakLineStrategy = breakline.EmptySpaceStrategy

	return text
}
