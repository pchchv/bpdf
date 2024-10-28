package error

import (
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/properties"
)

// DefaultErrorText is the default error text properties.
var DefaultErrorText = &properties.Text{
	Family: fontfamily.Arial,
	Style:  fontstyle.Bold,
	Size:   10,
	Color: &properties.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	},
}
