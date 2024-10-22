package fixture

import (
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/consts/breakline"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/consts/orientation"
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

// CellProp is responsible to give a valid properties.Cell.
func CellProp() properties.Cell {
	prop := properties.Cell{
		BackgroundColor: &properties.Color{
			Red:   255,
			Green: 100,
			Blue:  50,
		},
		BorderColor: &properties.Color{
			Red:   200,
			Green: 80,
			Blue:  60,
		},
		BorderType:      border.Left,
		BorderThickness: 0.6,
		LineStyle:       linestyle.Dashed,
	}
	return prop
}

// LineProp is responsible to give a valid properties.Line.
func LineProp() properties.Line {
	colorProp := ColorProp()
	prop := properties.Line{
		Color:         &colorProp,
		Style:         linestyle.Dashed,
		Thickness:     1.1,
		Orientation:   orientation.Vertical,
		OffsetPercent: 50,
		SizePercent:   20,
	}
	prop.MakeValid()
	return prop
}

// BarcodeProp is responsible to give a valid properties.Barcode.
func BarcodeProp() properties.Barcode {
	prop := properties.Barcode{
		Top:     10,
		Left:    10,
		Percent: 98,
		Proportion: properties.Proportion{
			Width:  16,
			Height: 9,
		},
		Center: false,
	}
	prop.MakeValid()
	return prop
}

// PageProp is responsible to give a valid properties.PageNumber.
func PageProp() properties.PageNumber {
	fontProp := FontProp()
	prop := properties.PageNumber{
		Pattern: "{current} / {total}",
		Place:   properties.LeftBottom,
		Family:  fontProp.Family,
		Style:   fontProp.Style,
		Size:    fontProp.Size,
		Color:   fontProp.Color,
	}
	return prop
}
