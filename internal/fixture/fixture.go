package fixture

import (
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/breakline"
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

// RectProp is responsible to give a valid properties.Rect.
func RectProp() properties.Rect {
	prop := properties.Rect{
		Top:     10,
		Left:    10,
		Percent: 98,
		Center:  false,
	}
	prop.MakeValid()
	return prop
}

// TextProp is responsible to give a valid properties.Text.
func TextProp() properties.Text {
	fontProp := FontProp()
	google := "https://www.google.com"
	prop := properties.Text{
		Top:               12,
		Bottom:            13,
		Left:              3,
		Family:            fontProp.Family,
		Style:             fontProp.Style,
		Size:              fontProp.Size,
		Align:             align.Right,
		BreakLineStrategy: breakline.DashStrategy,
		VerticalPadding:   20,
		Color:             fontProp.Color,
		Hyperlink:         &google,
	}
	prop.MakeValid(&fontProp)
	return prop
}
