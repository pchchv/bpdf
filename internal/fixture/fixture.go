package fixture

import (
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/properties"
)

// ColorProp is responsible to give a valid properties.Color.
func ColorProp() properties.Color {
	return properties.Color{
		Red:   100,
		Green: 50,
		Blue:  200,
	}
}

// FontProp is responsible to give a valid properties.Font.
func FontProp() properties.Font {
	colorProp := ColorProp()
	prop := properties.Font{
		Family: fontfamily.Helvetica,
		Style:  fontstyle.Bold,
		Size:   14,
		Color:  &colorProp,
	}
	prop.MakeValid(fontfamily.Arial)
	return prop
}
